package vrcosc

/*
Axes expect a float from -1 to 1 to control things like movement.

They expect to reset to 0 when not in use - otherwise a 'MoveForward' message left at '1' will continue to move you forward forever!

From https://docs.vrchat.com/docs/-as-input-controller
*/
type AxisInput string

const (
	// Move forwards (1) or Backwards (-1)
	Vertical AxisInput = "/input/Vertical"
	// Move right (1) or left (-1)
	Horizontal AxisInput = "/input/Horizontal"
	// Look Left and Right. Smooth in Desktop, VR will do a snap-turn when the value is 1 if Comfort Turning is on.
	CLookHorizontal AxisInput = "/input/LookHorizontal"
	// Use held item - not sure if this works
	UseAxisRight AxisInput = "/inpput/UseAxisRight"
	// Grab item - not sure if this works
	GrabAxisRight AxisInput = "/input/GrabAxisRight"
	// Move a held object forwards (1) and backwards (-1)
	MoveHoldFB AxisInput = "/input/MoveHoldFB"
	// Spin a held object Clockwise or Counter-Clockwise
	SpinHoldCwCcw AxisInput = "/input/SpinHoldCwCcw"
	// Spin a held object Up or Down
	SpinHoldUD AxisInput = "/input/SpinHoldUD"
	//  Spin a held object Left or Right
	SpinHoldLR AxisInput = "/input/SpinHoldLR"
)

/*
Buttons expect an int of 1 for 'pressed' and 0 for 'released'.

They will not function correctly without resetting to 0 first -
sending /input/Jump 1 and then /input/Jump 1 will only result in a single jump.

From https://docs.vrchat.com/docs/-as-input-controller
*/
type ButtonInput string

const (
	// Move forward while this is 1.
	MoveForward ButtonInput = "/input/MoveForward"
	// Move backwards while this is 1.
	MoveBackward ButtonInput = "/input/MoveBackward"
	// Strafe left while this is 1.
	MoveLeft ButtonInput = "/input/MoveLeft"
	// Strafe right while this is 1.
	MoveRight ButtonInput = "/input/MoveRight"
	// Turn to the left while this is 1. Smooth in Desktop, VR will do a snap-turn if Comfort Turning is on.
	LookLeft ButtonInput = "/input/LookLeft"
	// Turn to the right while this is 1. Smooth in Desktop, VR will do a snap-turn if Comfort Turning is on.
	LookRight ButtonInput = "/input/LookRight"
	// Jump if the world supports it.
	Jump ButtonInput = "/input/Jump"
	// Walk faster if the world supports it.
	Run ButtonInput = "/input/Run"
	// Snap-Turn to the left - VR Only.
	ComfortLeft ButtonInput = "/input/ComfortLeft"
	//  Snap-Turn to the right - VR Only.
	ComfortRight ButtonInput = "/input/ComfortRight"
	// Drop the item held in your right hand - VR Only.
	DropRight ButtonInput = "/input/DropRight"
	// Use the item highlighted by your right hand - VR Only.
	UseRight ButtonInput = "/input/UseRight"
	// Grab the item highlighted by your right hand - VR Only.
	GrabRight ButtonInput = "/input/GrabRight"
	// Drop the item held in your left hand - VR Only.
	DropLeft ButtonInput = "/input/DropLeft"
	// Use the item highlighted by your left hand - VR Only.
	UseLeft ButtonInput = "/input/UseLeft"
	// Grab the item highlighted by your left hand - VR Only.
	GrabLeft ButtonInput = "/input/GrabLeft"
	// Turn on Safe Mode.
	PanicButton ButtonInput = "/input/PanicButton"
	// Toggle QuickMenu On/Off. Will toggle upon receiving '1' if it's currently '0'.
	QuickMenuToggleLeft ButtonInput = "/input/QuickMenuToggleLeft"
	// Toggle QuickMenu On/Off. Will toggle upon receiving '1' if it's currently '0'.
	QuickMenuToggleRight ButtonInput = "/input/QuickMenuToggleRight"
	// Toggle Voice - the action will depend on whether "Toggle Voice" is turned on in your Settings. If so, then changing from 0 to 1 will toggle the state of mute. If "Toggle Voice" is turned off, then it functions like Push-To-Mute - 1 is muted, 0 is unmuted.
	Voice ButtonInput = "/input/Voice"
)

type ChatboxInput string

// Input text into the chatbox. If b is True, send the text in s immediately, bypassing the keyboard. If b is False, open the keyboard and populate it with the provided text.
const ChatInput ChatboxInput = "/chatbox/input"

// Toggle the typing indicator on or off.
const ChatTyping ChatboxInput = "/chatbox/typing"

type Avatar string

const (
	// When a new Avatar is loaded by the local player with  enabled, a message will be sent to /avatar/change with the ID of the avatar" - https://docs.vrchat.com/docs/-avatar-parameters
	AvatarChange Avatar = "/avatar/change"
)
