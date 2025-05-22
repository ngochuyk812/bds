import { create } from 'zustand';
import { Notification, NotificationState, NotificationType } from '../types/notification';

type NotificationStore = NotificationState & {
  addNotification: (notification: Omit<Notification, 'id' | 'createdAt'>) => string;
  success: (message: string, description?: string, duration?: number) => string;
  error: (message: string, description?: string, duration?: number) => string;
  errorExtras: (message: string, extras?: string[], duration?: number) => string;
  warning: (message: string, description?: string, duration?: number) => string;
  info: (message: string, description?: string, duration?: number) => string;
  removeNotification: (id: string) => void;
  clearNotifications: () => void;
};

const initialState: NotificationState = {
  notifications: [],
};

const generateId = (): string => Math.random().toString(36).substring(2, 9);

export const useNotificationStore = create<NotificationStore>((set, get) => ({
  ...initialState,

  addNotification: (notification) => {
    const id = generateId();
    const newNotification: Notification = {
      ...notification,
      id,
      createdAt: Date.now(),
    };
    if (!newNotification.description) {
      newNotification.description = newNotification.message;
      switch (newNotification.type) {
        case 'success':
          newNotification.description = 'Thành công';
          break;
        case 'error':
          newNotification.description = 'Lỗi';
          break;
        case 'warning':
          newNotification.description = 'Cảnh báo';
          break;
        case 'info':
          newNotification.description = 'Thông báo';
          break;
      }
    }

    set((state) => ({
      notifications: [...state.notifications, newNotification],
    }));

    if (notification.duration && notification.duration > 0) {
      setTimeout(() => {
        const stillExists = get().notifications.some((n) => n.id === id);
        if (!stillExists) return;

        set((state) => ({
          notifications: state.notifications.filter((n) => n.id !== id),
        }));

        if (notification.onClose) {
          notification.onClose();
        }
      }, notification.duration);
    }

    return id;
  },

  success: (message, description, duration = 3000) =>
    get().addNotification({ type: 'success', message, description, duration }),

  error: (message, description, duration = 5000) =>
    get().addNotification({ type: 'error', message, description, duration }),

  errorExtras: (message, extras, duration = 5000) =>
    get().addNotification({ type: 'error', message, description: extras?.join('\n'), duration }),

  warning: (message, description, duration = 4000) =>
    get().addNotification({ type: 'warning', message, description, duration }),

  info: (message, description, duration = 3000) =>
    get().addNotification({ type: 'info', message, description, duration }),

  removeNotification: (id) => {
    set((state) => ({
      notifications: state.notifications.filter((n) => n.id !== id),
    }));
  },

  clearNotifications: () => {
    set({ notifications: [] });
  },
}));
