package main

import "fmt"

type Device interface {
	IsEnabled() bool
	Enable()
	Disable()
	GetVolume() int
	SetVolume(volume int)
	GetChannel() int
	SetChannel(channel int)
	PrintStatus()
}

type TV struct {
	enabled bool
	volume  int
	channel int
}

func NewTV() *TV {
	return &TV{
		enabled: false,
		volume:  50,
		channel: 1,
	}
}

func (t *TV) IsEnabled() bool {
	return t.enabled
}

func (t *TV) Enable() {
	t.enabled = true
	fmt.Println("[TV] TV is now ON")
}

func (t *TV) Disable() {
	t.enabled = false
	fmt.Println("[TV] TV is now OFF")
}

func (t *TV) GetVolume() int {
	return t.volume
}

func (t *TV) SetVolume(volume int) {
	if volume > 100 {
		volume = 100
	} else if volume < 0 {
		volume = 0
	}
	t.volume = volume
	fmt.Printf("[TV] Volume set to %d%%\n", volume)
}

func (t *TV) GetChannel() int {
	return t.channel
}

func (t *TV) SetChannel(channel int) {
	t.channel = channel
	fmt.Printf("[TV] Channel set to %d\n", channel)
}

func (t *TV) PrintStatus() {
	status := "OFF"
	if t.enabled {
		status = "ON"
	}
	fmt.Printf("[TV] Status: %s | Volume: %d%% | Channel: %d\n", status, t.volume, t.channel)
}

type Radio struct {
	enabled   bool
	volume    int
	frequency float64
}

func NewRadio() *Radio {
	return &Radio{
		enabled:   false,
		volume:    30,
		frequency: 88.5,
	}
}

func (r *Radio) IsEnabled() bool {
	return r.enabled
}

func (r *Radio) Enable() {
	r.enabled = true
	fmt.Println("[Radio] Radio is now ON")
}

func (r *Radio) Disable() {
	r.enabled = false
	fmt.Println("[Radio] Radio is now OFF")
}

func (r *Radio) GetVolume() int {
	return r.volume
}

func (r *Radio) SetVolume(volume int) {
	if volume > 100 {
		volume = 100
	} else if volume < 0 {
		volume = 0
	}
	r.volume = volume
	fmt.Printf("[Radio] Volume set to %d%%\n", volume)
}

func (r *Radio) GetChannel() int {
	return int(r.frequency * 10)
}

func (r *Radio) SetChannel(channel int) {
	r.frequency = float64(channel) / 10.0
	fmt.Printf("[Radio] Frequency set to %.1f FM\n", r.frequency)
}

func (r *Radio) PrintStatus() {
	status := "OFF"
	if r.enabled {
		status = "ON"
	}
	fmt.Printf("[Radio] Status: %s | Volume: %d%% | Frequency: %.1f FM\n", status, r.volume, r.frequency)
}

type Remote struct {
	device Device
}

func NewRemote(device Device) *Remote {
	return &Remote{device: device}
}

func (r *Remote) TogglePower() {
	fmt.Println("\n[Remote] Toggling power...")
	if r.device.IsEnabled() {
		r.device.Disable()
	} else {
		r.device.Enable()
	}
}

func (r *Remote) VolumeUp() {
	fmt.Println("\n[Remote] Volume up...")
	currentVolume := r.device.GetVolume()
	r.device.SetVolume(currentVolume + 10)
}

func (r *Remote) VolumeDown() {
	fmt.Println("\n[Remote] Volume down...")
	currentVolume := r.device.GetVolume()
	r.device.SetVolume(currentVolume - 10)
}

func (r *Remote) ChannelUp() {
	fmt.Println("\n[Remote] Channel up...")
	currentChannel := r.device.GetChannel()
	r.device.SetChannel(currentChannel + 1)
}

func (r *Remote) ChannelDown() {
	fmt.Println("\n[Remote] Channel down...")
	currentChannel := r.device.GetChannel()
	r.device.SetChannel(currentChannel - 1)
}

func (r *Remote) Status() {
	fmt.Println("\n[Remote] Checking device status...")
	r.device.PrintStatus()
}

type AdvancedRemote struct {
	*Remote
}

func NewAdvancedRemote(device Device) *AdvancedRemote {
	return &AdvancedRemote{
		Remote: NewRemote(device),
	}
}

func (a *AdvancedRemote) Mute() {
	fmt.Println("\n[AdvancedRemote] Muting device...")
	a.device.SetVolume(0)
}

func (a *AdvancedRemote) SetChannel(channel int) {
	fmt.Printf("\n[AdvancedRemote] Setting channel to %d...\n", channel)
	a.device.SetChannel(channel)
}

type MessageSender interface {
	SendMessage(message string) error
}

type EmailSender struct {
	serverAddress string
}

func NewEmailSender(serverAddress string) *EmailSender {
	return &EmailSender{serverAddress: serverAddress}
}

func (e *EmailSender) SendMessage(message string) error {
	fmt.Printf("[Email] Sending via %s: %s\n", e.serverAddress, message)
	return nil
}

type SMSSender struct {
	gateway string
}

func NewSMSSender(gateway string) *SMSSender {
	return &SMSSender{gateway: gateway}
}

func (s *SMSSender) SendMessage(message string) error {
	fmt.Printf("[SMS] Sending via %s: %s\n", s.gateway, message)
	return nil
}

type PushSender struct {
	service string
}

func NewPushSender(service string) *PushSender {
	return &PushSender{service: service}
}

func (p *PushSender) SendMessage(message string) error {
	fmt.Printf("[Push] Sending via %s: %s\n", p.service, message)
	return nil
}

type Notification struct {
	sender MessageSender
}

func NewNotification(sender MessageSender) *Notification {
	return &Notification{sender: sender}
}

func (n *Notification) Send(message string) error {
	return n.sender.SendMessage(message)
}

type UrgentNotification struct {
	*Notification
}

func NewUrgentNotification(sender MessageSender) *UrgentNotification {
	return &UrgentNotification{
		Notification: NewNotification(sender),
	}
}

func (u *UrgentNotification) Send(message string) error {
	urgentMessage := "[URGENT] " + message
	return u.sender.SendMessage(urgentMessage)
}

type ScheduledNotification struct {
	*Notification
	schedule string
}

func NewScheduledNotification(sender MessageSender, schedule string) *ScheduledNotification {
	return &ScheduledNotification{
		Notification: NewNotification(sender),
		schedule:     schedule,
	}
}

func (s *ScheduledNotification) Send(message string) error {
	scheduledMessage := fmt.Sprintf("[Scheduled for %s] %s", s.schedule, message)
	return s.sender.SendMessage(scheduledMessage)
}

func main() {
	fmt.Println("=== Bridge Pattern Demo ===\n")

	fmt.Println("--- Example 1: Remote Control and Devices ---")
	
	tv := NewTV()
	basicRemote := NewRemote(tv)
	
	fmt.Println("\n** Using Basic Remote with TV **")
	basicRemote.Status()
	basicRemote.TogglePower()
	basicRemote.VolumeUp()
	basicRemote.VolumeUp()
	basicRemote.ChannelUp()
	basicRemote.Status()
	
	radio := NewRadio()
	advancedRemote := NewAdvancedRemote(radio)
	
	fmt.Println("\n** Using Advanced Remote with Radio **")
	advancedRemote.Status()
	advancedRemote.TogglePower()
	advancedRemote.VolumeUp()
	advancedRemote.SetChannel(1015)
	advancedRemote.Mute()
	advancedRemote.Status()
	
	fmt.Println("\n** Switching Advanced Remote to TV **")
	advancedRemoteTv := NewAdvancedRemote(tv)
	advancedRemoteTv.TogglePower()
	advancedRemoteTv.SetChannel(42)
	advancedRemoteTv.Status()

	fmt.Println("\n--- Example 2: Notification System ---")
	
	fmt.Println("\n** Regular Notifications **")
	emailSender := NewEmailSender("smtp.example.com")
	emailNotif := NewNotification(emailSender)
	emailNotif.Send("Welcome to our service!")
	
	smsSender := NewSMSSender("twilio.com")
	smsNotif := NewNotification(smsSender)
	smsNotif.Send("Your verification code is 123456")
	
	pushSender := NewPushSender("firebase.com")
	pushNotif := NewNotification(pushSender)
	pushNotif.Send("You have a new message")
	
	fmt.Println("\n** Urgent Notifications **")
	urgentEmail := NewUrgentNotification(emailSender)
	urgentEmail.Send("Server is down!")
	
	urgentSMS := NewUrgentNotification(smsSender)
	urgentSMS.Send("Security alert detected")
	
	fmt.Println("\n** Scheduled Notifications **")
	scheduledEmail := NewScheduledNotification(emailSender, "2024-12-25 09:00")
	scheduledEmail.Send("Merry Christmas!")
	
	scheduledPush := NewScheduledNotification(pushSender, "2024-01-01 00:00")
	scheduledPush.Send("Happy New Year!")

	fmt.Println("\n--- Summary ---")
	fmt.Println("Bridge pattern separates abstraction from implementation")
	fmt.Println("Remotes (abstraction) work with any device (implementation)")
	fmt.Println("Notifications (abstraction) work with any sender (implementation)")
	fmt.Println("Both hierarchies can evolve independently")
}
