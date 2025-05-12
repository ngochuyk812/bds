import React from 'react';
import { useNotificationStore } from '../store/notification';
import NotificationItem from './NotificationItem';

const NotificationContainer: React.FC = () => {
  const { notifications } = useNotificationStore();

  if (notifications.length === 0) {
    return null;
  }

  return (
    <div className="fixed top-4 right-4 flex flex-col items-end z-[9999]">
      {notifications.map((notification) => (
        <NotificationItem key={notification.id} notification={notification} />
      ))}
    </div>
  );
};

export default NotificationContainer;
