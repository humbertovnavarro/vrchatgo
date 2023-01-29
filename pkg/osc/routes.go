package osc

/*
Axes expect a float from -1 to 1 to control things like movement.

They expect to reset to 0 when not in use - otherwise a 'MoveForward' message left at '1' will continue to move you forward forever!

From https://docs.vrchat.com/docs/osc-as-input-controller
*/
type OSCAxisInput string

const (
	// Move forwards (1) or Backwards (-1)
	OSCVertical OSCAxisInput = "/input/Vertical"
	// Move right (1) or left (-1)
	OSCHorizontal OSCAxisInput = "/input/Horizontal"
	// Look Left and Right. Smooth in Desktop, VR will do a snap-turn when the value is 1 if Comfort Turning is on.
	OSCCLookHorizontal OSCAxisInput = "/input/LookHorizontal"
	// Use held item - not sure if this works
	OSCUseAxisRight OSCAxisInput = "/inpput/UseAxisRight"
	// Grab item - not sure if this works
	OSCGrabAxisRight OSCAxisInput = "/input/GrabAxisRight"
	// Move a held object forwards (1) and backwards (-1)
	OSCMoveHoldFB OSCAxisInput = "/input/MoveHoldFB"
	// Spin a held object Clockwise or Counter-Clockwise
	OSCSpinHoldCwCcw OSCAxisInput = "/input/SpinHoldCwCcw"
	// Spin a held object Up or Down
	OSCSpinHoldUD OSCAxisInput = "/input/SpinHoldUD"
	//  Spin a held object Left or Right
	OSCSpinHoldLR OSCAxisInput = "/input/SpinHoldLR"
)

/*
Buttons expect an int of 1 for 'pressed' and 0 for 'released'.

They will not function correctly without resetting to 0 first -
sending /input/Jump 1 and then /input/Jump 1 will only result in a single jump.

From https://docs.vrchat.com/docs/osc-as-input-controller
*/
type OSCButtonInput string

const (
	// Move forward while this is 1.
	OSCMoveForward OSCButtonInput = "/input/MoveForward"
	// Move backwards while this is 1.
	OSCMoveBackward OSCButtonInput = "/input/MoveBackward"
	// Strafe left while this is 1.
	OSCMoveLeft OSCButtonInput = "/input/MoveLeft"
	// Strafe right while this is 1.
	OSCMoveRight OSCButtonInput = "/input/MoveRight"
	// Turn to the left while this is 1. Smooth in Desktop, VR will do a snap-turn if Comfort Turning is on.
	OSCLookLeft OSCButtonInput = "/input/LookLeft"
	// Turn to the right while this is 1. Smooth in Desktop, VR will do a snap-turn if Comfort Turning is on.
	OSCLookRight OSCButtonInput = "/input/LookRight"
	// Jump if the world supports it.
	OSCJump OSCButtonInput = "/input/Jump"
	// Walk faster if the world supports it.
	OSCRun OSCButtonInput = "/input/Run"
	// Snap-Turn to the left - VR Only.
	OSCComfortLeft OSCButtonInput = "/input/ComfortLeft"
	//  Snap-Turn to the right - VR Only.
	OSCComfortRight OSCButtonInput = "/input/ComfortRight"
	// Drop the item held in your right hand - VR Only.
	OSCDropRight OSCButtonInput = "/input/DropRight"
	// Use the item highlighted by your right hand - VR Only.
	OSCUseRight OSCButtonInput = "/input/UseRight"
	// Grab the item highlighted by your right hand - VR Only.
	OSCGrabRight OSCButtonInput = "/input/GrabRight"
	// Drop the item held in your left hand - VR Only.
	OSCDropLeft OSCButtonInput = "/input/DropLeft"
	// Use the item highlighted by your left hand - VR Only.
	OSCUseLeft OSCButtonInput = "/input/UseLeft"
	// Grab the item highlighted by your left hand - VR Only.
	OSCGrabLeft OSCButtonInput = "/input/GrabLeft"
	// Turn on Safe Mode.
	OSCPanicButton OSCButtonInput = "/input/PanicButton"
	// Toggle QuickMenu On/Off. Will toggle upon receiving '1' if it's currently '0'.
	OSCQuickMenuToggleLeft OSCButtonInput = "/input/QuickMenuToggleLeft"
	// Toggle QuickMenu On/Off. Will toggle upon receiving '1' if it's currently '0'.
	OSCQuickMenuToggleRight OSCButtonInput = "/input/QuickMenuToggleRight"
	// Toggle Voice - the action will depend on whether "Toggle Voice" is turned on in your Settings. If so, then changing from 0 to 1 will toggle the state of mute. If "Toggle Voice" is turned off, then it functions like Push-To-Mute - 1 is muted, 0 is unmuted.
	OSCVoice OSCButtonInput = "/input/Voice"
)

type OSCChatboxInput string

// Input text into the chatbox. If b is True, send the text in s immediately, bypassing the keyboard. If b is False, open the keyboard and populate it with the provided text.
const OSCChatInput OSCChatboxInput = "/chatbox/input"

// Toggle the typing indicator on or off.
const OSCChatTyping OSCChatboxInput = "/chatbox/typing"

type OSCAvatar string

const (
	// When a new Avatar is loaded by the local player with OSC enabled, a message will be sent to /avatar/change with the ID of the avatar" - https://docs.vrchat.com/docs/osc-avatar-parameters
	OSCAvatarChange OSCAvatar = "/avatar/change"
)
