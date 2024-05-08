package api

import (
	"fmt"
	"time"
)

// ButtonStyle is a style of Button element.
type ButtonStyle string

// Represents a general button styles.
const (
	ButtonStyleDefault ButtonStyle = ""
	ButtonStylePrimary ButtonStyle = "primary"
	ButtonStyleDanger  ButtonStyle = "danger"
)

// SelectType is a type of Button element.
type SelectType string

// Represents a select dropdown types.
const (
	StaticSelect   SelectType = "static"
	ExternalSelect SelectType = "external"
)

// DividerStyle is a style of Divider element between section blocks.
type DividerStyle string

// Represents a divider styles.
const (
	// DividerStyleDefault put a block divider, like an <hr>, to split up different sections inside of a single message.
	// It is the default style for backwards compatibility.
	DividerStyleDefault DividerStyle = ""
	// DividerStyleTopNone
	DividerStyleTopNone DividerStyle = "none"
)

// MessageType defines the message type.
type MessageType string

const (
	// DefaultMessage defines a message that should be displayed in default mode supported by communicator.
	DefaultMessage MessageType = ""
	// BasicCardWithButtonsInSeparateMessage defines a message that should be displayed in plaintext mode supported by the communicator,
	// with the buttons sent in a separate interactive message.
	// This feature is currently available only for the Teams platform.
	BasicCardWithButtonsInSeparateMsg = "basicCardWithButtonsInSeparateMessage"
	// BaseBodyWithFilterMessage defines a message that should be displayed in plaintext mode supported by communicator.
	// In this form the built-in filter is supported.
	// NOTE: only BaseBody is preserved. All other properties are ignored even if set.
	BaseBodyWithFilterMessage MessageType = "baseBodyWithFilter"
	// NonInteractiveSingleSection it is an indicator for non-interactive platforms, that they can render this event
	// even though they have limited capability. As a result, a given message has the following restriction:
	//  - the whole message should have exactly one section
	//  - section interactive elements such as buttons, select, multiselect, and inputs are ignored.
	//  - the base body of the message is ignored
	//  - Timestamp field is optional
	NonInteractiveSingleSection MessageType = "nonInteractiveEventSingleSection"
	// PopupMessage defines a message that should be displayed to the user as popup (if possible).
	PopupMessage MessageType = "form"
	// ThreadMessage defines a message that should be sent in a thread.
	ThreadMessage MessageType = "threadMessage"
	// SkipMessage defines a message that should not be sent to the end user.
	// If not used and message is empty, a special indicator will be sent as the response.
	SkipMessage MessageType = "skipMessage"
)

// Message represents a generic message with interactive buttons.
type Message struct {
	Type              MessageType `json:"type,omitempty" yaml:"type"`
	BaseBody          Body        `json:"baseBody,omitempty" yaml:"baseBody"`
	Timestamp         time.Time   `json:"timestamp,omitempty" yaml:"timestamp"`
	Sections          []Section   `json:"sections,omitempty" yaml:"sections"`
	PlaintextInputs   LabelInputs `json:"plaintextInputs,omitempty" yaml:"plaintextInputs"`
	OnlyVisibleForYou bool        `json:"onlyVisibleForYou,omitempty" yaml:"onlyVisibleForYou"`
	ReplaceOriginal   bool        `json:"replaceOriginal,omitempty" yaml:"replaceOriginal"`
	UserHandle        string      `json:"userHandle,omitempty" yaml:"userHandle"`

	// ParentActivityID represents the originating message that started a thread. If set, message will be sent in that thread instead of the default one.
	ParentActivityID string `json:"parentActivityId,omitempty" yaml:"parentActivityId,omitempty"`
}

func (msg *Message) IsEmpty() bool {
	if msg.HasBaseBody() {
		return false
	}
	if msg.HasInputs() {
		return false
	}
	if msg.HasSections() {
		return false
	}
	if !msg.Timestamp.IsZero() {
		return false
	}

	return true
}

// HasBaseBody returns true if message has base body defined.
func (msg *Message) HasBaseBody() bool {
	var emptyBase Body
	return msg.BaseBody != emptyBase
}

// HasSections returns true if message has interactive sections.
func (msg *Message) HasSections() bool {
	return len(msg.Sections) != 0
}

// HasInputs returns true if message has interactive inputs.
func (msg *Message) HasInputs() bool {
	return len(msg.PlaintextInputs) != 0
}

// Select holds data related to the select drop-down.
type Select struct {
	Type    SelectType `json:"type,omitempty" yaml:"type"`
	Name    string     `json:"name,omitempty" yaml:"name"`
	Command string     `json:"command,omitempty" yaml:"command"`
	// OptionGroups provides a way to group options in a select menu.
	OptionGroups []OptionGroup `json:"optionGroups,omitempty" yaml:"optionGroups"`
	// InitialOption holds already pre-selected options. MUST be a sub-set of OptionGroups.
	InitialOption *OptionItem `json:"initialOption,omitempty" yaml:"initialOption"`
}

// Base holds generic message fields.
type Base struct {
	Header      string `json:"header,omitempty" yaml:"header"`
	Description string `json:"description,omitempty" yaml:"description"`
	Body        Body   `json:"body,omitempty" yaml:"body"`
}

// Body holds message body fields.
type Body struct {
	CodeBlock string `json:"codeBlock,omitempty" yaml:"codeBlock"`
	Plaintext string `json:"plaintext,omitempty" yaml:"plaintext"`
}

// SectionStyle holds section style.
type SectionStyle struct {
	Divider DividerStyle `json:"divider,omitempty" yaml:"dividerStyle"`
}

// Section holds section related fields.
type Section struct {
	Style SectionStyle `json:"style,omitempty" yaml:"style"`

	Base            `json:",inline" yaml:"base"`
	Buttons         Buttons      `json:"buttons,omitempty" yaml:"buttons"`
	MultiSelect     MultiSelect  `json:"multiSelect,omitempty" yaml:"multiSelect"`
	Selects         Selects      `json:"selects,omitempty" yaml:"selects"`
	PlaintextInputs LabelInputs  `json:"plaintextInputs,omitempty" yaml:"plaintextInputs"`
	TextFields      TextFields   `json:"textFields,omitempty" yaml:"textFields"`
	BulletLists     BulletLists  `json:"bulletLists,omitempty" yaml:"bulletLists"`
	Context         ContextItems `json:"context,omitempty" yaml:"context"`
}

// BulletLists holds the bullet lists.
type BulletLists []BulletList

// AreItemsDefined returns true if at least one list has items defined.
func (l BulletLists) AreItemsDefined() bool {
	for _, list := range l {
		if len(list.Items) > 0 {
			return true
		}
	}
	return false
}

// LabelInputs holds the plain text input items.
type LabelInputs []LabelInput

// ContextItems holds context items.
type ContextItems []ContextItem

// TextFields holds text field items.
type TextFields []TextField

// TextField holds a text field data.
type TextField struct {
	Key   string `json:"key,omitempty" yaml:"key"`
	Value string `json:"value,omitempty" yaml:"value"`
}

// IsEmpty returns true if all fields have zero-value.
func (t *TextField) IsEmpty() bool {
	return t.Value == "" && t.Key == ""
}

// BulletList defines a bullet list primitive.
type BulletList struct {
	Title string   `json:"title,omitempty" yaml:"title"`
	Items []string `json:"items,omitempty" yaml:"items"`
}

// IsDefined returns true if there are any context items defined.
func (c ContextItems) IsDefined() bool {
	return len(c) > 0
}

// ContextItem holds context item.
type ContextItem struct {
	Text string `json:"text,omitempty" yaml:"text"`
}

// Selects holds multiple Select objects.
type Selects struct {
	// ID allows to identify a given block when we do the updated.
	ID    string   `json:"id,omitempty" yaml:"id"`
	Items []Select `json:"items,omitempty" yaml:"items"`
}

// DispatchedInputAction defines when the action should be sent to our backend.
type DispatchedInputAction string

// Defines the possible options to dispatch the input action.
const (
	NoDispatchInputAction          DispatchedInputAction = ""
	DispatchInputActionOnEnter     DispatchedInputAction = "on_enter_pressed"
	DispatchInputActionOnCharacter DispatchedInputAction = "on_character_entered"
)

// LabelInput is used to create input elements to use in messages.
type LabelInput struct {
	Command          string                `json:"command,omitempty" yaml:"command"`
	Text             string                `json:"text,omitempty" yaml:"text"`
	Placeholder      string                `json:"placeholder,omitempty" yaml:"placeholder"`
	DispatchedAction DispatchedInputAction `json:"dispatchedAction,omitempty" yaml:"dispatchedAction"`
}

// AreOptionsDefined returns true if some options are available.
func (s *Selects) AreOptionsDefined() bool {
	if s == nil {
		return false
	}
	return len(s.Items) > 0
}

// OptionItem defines an option model.
type OptionItem struct {
	Name  string `json:"name,omitempty" yaml:"name"`
	Value string `json:"value,omitempty" yaml:"value"`
}

// MultiSelect holds multi select related fields.
type MultiSelect struct {
	Name        string `json:"name,omitempty" yaml:"name"`
	Description Body   `json:"description,omitempty" yaml:"description"`
	Command     string `json:"command,omitempty" yaml:"command"`

	// Options holds all available options
	Options []OptionItem `json:"options,omitempty" yaml:"options"`

	// InitialOptions hold already pre-selected options. MUST be a sub-set of Options.
	InitialOptions []OptionItem `json:"initialOptions,omitempty" yaml:"initialOptions"`
}

// OptionGroup holds information about options in the same group.
type OptionGroup struct {
	Name    string       `json:"name,omitempty" yaml:"name"`
	Options []OptionItem `json:"options,omitempty" yaml:"options"`
}

// AreOptionsDefined returns true if some options are available.
func (m *MultiSelect) AreOptionsDefined() bool {
	if m == nil {
		return false
	}
	if len(m.Options) == 0 {
		return false
	}
	return true
}

// Buttons holds definition of interactive buttons.
type Buttons []Button

// GetButtonsWithDescription returns all buttons with description.
func (s *Buttons) GetButtonsWithDescription() Buttons {
	if s == nil {
		return nil
	}
	var out Buttons
	for _, item := range *s {
		if item.Description == "" {
			continue
		}
		out = append(out, item)
	}

	return out
}

// GetButtonsWithoutDescription returns all buttons without description.
func (s *Buttons) GetButtonsWithoutDescription() Buttons {
	if s == nil {
		return nil
	}
	var out Buttons
	for _, item := range *s {
		if item.Description != "" {
			continue
		}
		out = append(out, item)
	}

	return out
}

// ButtonDescriptionStyle defines the style of the button description.
type ButtonDescriptionStyle string

const (
	// ButtonDescriptionStyleBold defines the bold style for the button description.
	ButtonDescriptionStyleBold ButtonDescriptionStyle = "bold"
	// ButtonDescriptionStyleText defines the plaintext style for the button description.
	ButtonDescriptionStyleText ButtonDescriptionStyle = "text"
	// ButtonDescriptionStyleCode defines the code style for the button description.
	ButtonDescriptionStyleCode ButtonDescriptionStyle = "code"
)

// Button holds definition of action button.
type Button struct {
	Description string `json:"description,omitempty" yaml:"description"`

	// DescriptionStyle defines the style of the button description. If not provided, the default style (ButtonDescriptionStyleCode) is used.
	DescriptionStyle ButtonDescriptionStyle `json:"descriptionStyle" yaml:"descriptionStyle"`

	Name    string      `json:"name,omitempty" yaml:"name"`
	Command string      `json:"command,omitempty" yaml:"command"`
	URL     string      `json:"url,omitempty" yaml:"url"`
	Style   ButtonStyle `json:"style,omitempty" yaml:"style"`
}

// ButtonBuilder provides a simplified way to construct a Button model.
type ButtonBuilder struct{}

func NewMessageButtonBuilder() *ButtonBuilder {
	return &ButtonBuilder{}
}

// ForCommandWithDescCmd returns button command where description and command are the same.
func (b *ButtonBuilder) ForCommandWithDescCmd(name, cmd string, style ...ButtonStyle) Button {
	bt := ButtonStyleDefault
	if len(style) > 0 {
		bt = style[0]
	}
	return b.commandWithCmdDesc(name, cmd, cmd, bt)
}

// ForCommandWithBoldDesc returns button command where description and command are different.
func (b *ButtonBuilder) ForCommandWithBoldDesc(name, desc, cmd string, style ...ButtonStyle) Button {
	bt := ButtonStyleDefault
	if len(style) > 0 {
		bt = style[0]
	}
	return b.commandWithDesc(name, cmd, desc, bt, ButtonDescriptionStyleBold)
}

// DescriptionURL returns link button with description.
func (b *ButtonBuilder) DescriptionURL(name, cmd string, url string, style ...ButtonStyle) Button {
	bt := ButtonStyleDefault
	if len(style) > 0 {
		bt = style[0]
	}

	return Button{
		Name:        name,
		Description: fmt.Sprintf("%s %s", MessageBotNamePlaceholder, cmd),
		URL:         url,
		Style:       bt,
	}
}

// ForCommandWithoutDesc returns button command without description.
func (b *ButtonBuilder) ForCommandWithoutDesc(name, cmd string, style ...ButtonStyle) Button {
	bt := ButtonStyleDefault
	if len(style) > 0 {
		bt = style[0]
	}
	cmd = fmt.Sprintf("%s %s", MessageBotNamePlaceholder, cmd)
	return Button{
		Name:    name,
		Command: cmd,
		Style:   bt,
	}
}

// ForCommand returns button command with description in adaptive code block.
//
// For displaying description in bold, use ForCommandWithBoldDesc.
func (b *ButtonBuilder) ForCommand(name, cmd, desc string, style ...ButtonStyle) Button {
	bt := ButtonStyleDefault
	if len(style) > 0 {
		bt = style[0]
	}
	return b.commandWithCmdDesc(name, cmd, desc, bt)
}

// ForURLWithBoldDesc returns link button with description.
func (b *ButtonBuilder) ForURLWithBoldDesc(name, desc, url string, style ...ButtonStyle) Button {
	urlBtn := b.ForURL(name, url, style...)
	urlBtn.Description = desc
	urlBtn.DescriptionStyle = ButtonDescriptionStyleBold

	return urlBtn
}

// ForURLWithBoldDesc returns link button with description.
func (b *ButtonBuilder) ForURLWithTextDesc(name, desc, url string, style ...ButtonStyle) Button {
	urlBtn := b.ForURL(name, url, style...)
	urlBtn.Description = desc
	urlBtn.DescriptionStyle = ButtonDescriptionStyleText

	return urlBtn
}

// ForURL returns link button.
func (b *ButtonBuilder) ForURL(name, url string, style ...ButtonStyle) Button {
	bt := ButtonStyleDefault
	if len(style) > 0 {
		bt = style[0]
	}

	return Button{
		Name:  name,
		URL:   url,
		Style: bt,
	}
}

func (b *ButtonBuilder) commandWithCmdDesc(name, cmd, desc string, style ButtonStyle) Button {
	desc = fmt.Sprintf("%s %s", MessageBotNamePlaceholder, desc)
	return b.commandWithDesc(name, cmd, desc, style, ButtonDescriptionStyleCode)
}

func (b *ButtonBuilder) commandWithDesc(name, cmd, desc string, style ButtonStyle, descStyle ButtonDescriptionStyle) Button {
	cmd = fmt.Sprintf("%s %s", MessageBotNamePlaceholder, cmd)
	return Button{
		Name:             name,
		Command:          cmd,
		Description:      desc,
		DescriptionStyle: descStyle,
		Style:            style,
	}
}
