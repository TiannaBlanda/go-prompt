// Code generated by "stringer -type=Key"; DO NOT EDIT.

package prompt

import "strconv"

const _Key_name = "EscapeControlAControlBControlCControlDControlEControlFControlGControlHControlIControlJControlKControlLControlMControlNControlOControlPControlQControlRControlSControlTControlUControlVControlWControlXControlYControlZControlSpaceControlBackslashControlSquareCloseControlCircumflexControlUnderscoreControlLeftControlRightControlUpControlDownUpDownRightLeftShiftLeftShiftUpShiftDownShiftRightHomeEndDeleteShiftDeleteControlDeletePageUpPageDownBackTabInsertBackspaceTabEnterF1F2F3F4F5F6F7F8F9F10F11F12F13F14F15F16F17F18F19F20F21F22F23F24AnyCPRResponseVt100MouseEventWindowsMouseEventBracketedPasteIgnoreNotDefined"

var _Key_index = [...]uint16{0, 6, 14, 22, 30, 38, 46, 54, 62, 70, 78, 86, 94, 102, 110, 118, 126, 134, 142, 150, 158, 166, 174, 182, 190, 198, 206, 214, 226, 242, 260, 277, 294, 305, 317, 326, 337, 339, 343, 348, 352, 361, 368, 377, 387, 391, 394, 400, 411, 424, 430, 438, 445, 451, 460, 463, 468, 470, 472, 474, 476, 478, 480, 482, 484, 486, 489, 492, 495, 498, 501, 504, 507, 510, 513, 516, 519, 522, 525, 528, 531, 534, 545, 560, 577, 591, 597, 607}

func (i Key) String() string {
	if i < 0 || i >= Key(len(_Key_index)-1) {
		return "Key(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Key_name[_Key_index[i]:_Key_index[i+1]]
}
 