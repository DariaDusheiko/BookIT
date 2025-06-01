import dayjs from 'dayjs';
import 'dayjs/locale/ru';

// Set Russian locale
dayjs.locale('ru');

export const formatDateTime = (dateTime: string) => {
  return dayjs(dateTime).format('D MMMM YYYY HH:mm');
}; 