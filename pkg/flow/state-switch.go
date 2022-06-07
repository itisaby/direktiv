package flow

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/direktiv/direktiv/pkg/model"
)

type switchStateLogic struct {
	*model.SwitchState
}

func initSwitchStateLogic(wf *model.Workflow, state model.State) (stateLogic, error) {

	switchState, ok := state.(*model.SwitchState)
	if !ok {
		return nil, NewInternalError(errors.New("bad state object"))
	}

	sl := new(switchStateLogic)
	sl.SwitchState = switchState

	return sl, nil
}

func (sl *switchStateLogic) Deadline(ctx context.Context, engine *engine, im *instanceMemory) time.Time {
	return time.Now().Add(defaultDeadline)
}

func (sl *switchStateLogic) Run(ctx context.Context, engine *engine, im *instanceMemory, wakedata []byte) (transition *stateTransition, err error) {

	if im.GetMemory() != nil {
		err = NewInternalError(errors.New("got unexpected savedata"))
		return
	}

	if len(wakedata) != 0 {
		err = NewInternalError(errors.New("got unexpected wakedata"))
		return
	}

	for i, condition := range sl.Conditions {

		var x interface{}
		x, err = jqOne(im.data, condition.Condition)
		if err != nil {
			err = NewInternalError(fmt.Errorf("switch condition %d condition failed to run: %v", i, err))
			return
		}

		if truth(x) {
			engine.logToInstance(ctx, time.Now(), im.in, "Switch condition %d succeeded", i)
			transition = &stateTransition{
				Transform: condition.Transform,
				NextState: condition.Transition,
			}
			return
		}

	}

	engine.logToInstance(ctx, time.Now(), im.in, "No switch conditions succeeded")
	transition = &stateTransition{
		Transform: sl.DefaultTransform,
		NextState: sl.DefaultTransition,
	}

	return

}
