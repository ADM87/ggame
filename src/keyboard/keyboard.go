package keyboard

import (
	"github.com/ADM87/ggame/src/sys"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type KeyPhase int
type KeyAction func() error
type KeyPhaseActions map[ebiten.Key]map[KeyPhase]KeyAction

func (phase KeyPhase) String() string {
	switch phase {
	case KeyPhaseDown:
		return "down"
	case KeyPhaseHeld:
		return "held"
	case KeyPhaseUp:
		return "up"
	case KeyPhaseNone:
		return "none"
	default:
		return "unknown"
	}
}

const (
	KeyPhaseDown KeyPhase = iota
	KeyPhaseHeld
	KeyPhaseUp
	KeyPhaseNone
)

var (
	keyRegistry = make(KeyPhaseActions)
)

func RegisterKey(key ebiten.Key, phase KeyPhase, action KeyAction) {
	if _, exists := keyRegistry[key]; !exists {
		keyRegistry[key] = make(map[KeyPhase]KeyAction)
	}

	if _, exists := keyRegistry[key][phase]; !exists {
		sys.Logger().Debugf("Registering key action for key %s and phase %s", key, phase)
	} else {
		sys.Logger().Warnf("Key action for key %s and phase %s already exists, overwriting", key, phase)
	}

	keyRegistry[key][phase] = action
}

func UnregisterKey(key ebiten.Key) {
	if _, exists := keyRegistry[key]; exists {
		sys.Logger().Debugf("Unregistering all actions for key %s", key)
		delete(keyRegistry, key)
	} else {
		sys.Logger().Warnf("No actions registered for key %s, nothing to unregister", key)
	}
}

func Ping() error {
	for key, actions := range keyRegistry {
		switch {
		case inpututil.IsKeyJustPressed(key):
			if action, exists := actions[KeyPhaseDown]; exists {
				sys.Logger().Debugf("Key %s pressed, executing action for phase %s", key, KeyPhaseDown)
				if err := action(); err != nil {
					return err
				}
			}
		case inpututil.IsKeyJustReleased(key):
			if action, exists := actions[KeyPhaseUp]; exists {
				sys.Logger().Debugf("Key %s released, executing action for phase %s", key, KeyPhaseUp)
				if err := action(); err != nil {
					return err
				}
			}
		case ebiten.IsKeyPressed(key):
			if action, exists := actions[KeyPhaseHeld]; exists {
				sys.Logger().Debugf("Key %s is held down, executing action for phase %s", key, KeyPhaseHeld)
				if err := action(); err != nil {
					return err
				}
			}
		}

		if len(keyRegistry) == 0 {
			sys.Logger().Warnf("Key %s purged key phase actions, unable to continue processing", key)
			return nil
		}
	}

	return nil
}
