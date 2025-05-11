export type NotificationType = 'success' | 'error' | 'warning' | 'info';

export interface Notification {
  id: string;
  type: NotificationType;
  message: string;
  description?: string;
  duration?: number;
  onClose?: () => void;
  createdAt: number;
}

export interface NotificationState {
  notifications: Notification[];
}
