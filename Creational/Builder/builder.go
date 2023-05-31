package main

import "fmt"

type NotificationBuilder struct {
	Title    string
	SubTitle string
	Message  string
	Icon     string
	Image    string
	Priority int
	NotType  string
}

func newNotificationBuilder() *NotificationBuilder {
	return &NotificationBuilder{}
}

func (nb *NotificationBuilder) SetTitle(title string) {
	nb.Title = title
}

func (nb *NotificationBuilder) SetSubTitle(subTitle string) {
	nb.SubTitle = subTitle
}

func (nb *NotificationBuilder) SetMessage(message string) {
	nb.Message = message
}

func (nb *NotificationBuilder) SetImage(image string) {
	nb.Image = image
}

func (nb *NotificationBuilder) SetIcon(icon string) {
	nb.Icon = icon
}

func (nb *NotificationBuilder) SetPriority(pri int) {
	nb.Priority = pri
}

func (nb *NotificationBuilder) SetType(notType string) {
	nb.NotType = notType
}

func (nb *NotificationBuilder) Build() (*Notification, error) {

	if nb.Priority > 5 {
		return nil, fmt.Errorf("priority can't be greater than 5")
	}

	return &Notification{
		title:    nb.Title,
		subtitle: nb.SubTitle,
		message:  nb.Message,
		icon:     nb.Icon,
		image:    nb.Image,
		priority: nb.Priority,
		notType:  nb.NotType,
	}, nil
}
