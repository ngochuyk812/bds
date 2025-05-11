import React, { useEffect, useState } from 'react';
import { CheckCircleOutlined, CloseCircleOutlined, InfoCircleOutlined, WarningOutlined, CloseOutlined } from '@ant-design/icons';
import { Notification } from '../types/notification';
import { useNotificationStore } from '../store/notification';

interface NotificationItemProps {
  notification: Notification;
}

const NotificationItem: React.FC<NotificationItemProps> = ({ notification }) => {
  const { removeNotification } = useNotificationStore();
  const [isExiting, setIsExiting] = useState(false);

  const handleClose = () => {
    setIsExiting(true);
    setTimeout(() => {
      removeNotification(notification.id);
      if (notification.onClose) {
        notification.onClose();
      }
    }, 300);
  };

  useEffect(() => {
    if (notification.duration && notification.duration > 0) {
      const timer = setTimeout(() => {
        handleClose();
      }, notification.duration);
      return () => clearTimeout(timer);
    }
  }, [notification.id, notification.duration]);

  const getIcon = () => {
    switch (notification.type) {
      case 'success':
        return <CheckCircleOutlined className="text-green-500 text-xl" />;
      case 'error':
        return <CloseCircleOutlined className="text-red-500 text-xl" />;
      case 'warning':
        return <WarningOutlined className="text-yellow-500 text-xl" />;
      case 'info':
        return <InfoCircleOutlined className="text-blue-500 text-xl" />;
      default:
        return null;
    }
  };

  const getBackgroundColor = () => {
    switch (notification.type) {
      case 'success':
        return 'bg-green-50 border-green-200';
      case 'error':
        return 'bg-red-50 border-red-200';
      case 'warning':
        return 'bg-yellow-50 border-yellow-200';
      case 'info':
        return 'bg-blue-50 border-blue-200';
      default:
        return 'bg-white';
    }
  };

  return (
    <div
      className={`${getBackgroundColor()} border rounded-md shadow-md p-4 mb-2 flex items-start transition-all duration-300 ${isExiting ? 'opacity-0 translate-x-full' : 'opacity-100'
        }`}
      style={{ minWidth: '300px', maxWidth: '400px' }}
    >
      <div className="mr-3 mt-1">{getIcon()}</div>
      <div className="flex-1">
        <div className="font-semibold">{notification.message}</div>
        {notification.description && <div className="text-sm mt-1">{notification.description}</div>}
      </div>
      <button
        onClick={handleClose}
        className="ml-2 text-gray-400 hover:text-gray-600 focus:outline-none"
        aria-label="Close notification"
      >
        <CloseOutlined />
      </button>
    </div>
  );
};

export default NotificationItem;
