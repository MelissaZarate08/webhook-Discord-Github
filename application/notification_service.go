package application

import (
    "fmt"
    "webhook-github/domain"
    discord "webhook-github/infrastructure/adapters/outbound"
)


type NotificationService struct {
	discordNotifier discord.Notifier
}


func NewNotificationService(notifier discord.Notifier) *NotificationService {
	return &NotificationService{
		discordNotifier: notifier,
	}
}


func (s *NotificationService) NotifyPullRequestEvent(evt domain.PullRequestEvent) error {
	var message string
	switch evt.Action {
	case "opened":
		message = fmt.Sprintf("Nuevo PR: **%s** (#%d) abierto por @%s", evt.Title, evt.Number, evt.User)
	case "reopened":
		message = fmt.Sprintf("PR reabierto: **%s** (#%d)", evt.Title, evt.Number)
	case "ready_for_review":
		message = fmt.Sprintf("PR listo para revisión: **%s** (#%d)", evt.Title, evt.Number)
	case "closed":
		message = fmt.Sprintf("PR cerrado o fusionado: **%s** (#%d)", evt.Title, evt.Number)
	default:
		return nil 
	}

	return s.discordNotifier.Send(message)
}


func (s *NotificationService) NotifyActionsEvent(evt domain.ActionsEvent) error {
	var message string
	switch evt.Conclusion {
	case "success":
		message = fmt.Sprintf("El workflow **%s** ha pasado exitosamente.", evt.Workflow)
	case "failure":
		message = fmt.Sprintf("El workflow **%s** ha fallado.", evt.Workflow)
	default:
		message = fmt.Sprintf("Estado desconocido del workflow **%s**: %s", evt.Workflow, evt.Conclusion)
	}

	return s.discordNotifier.Send(message)
}
