export type searchTypeI = 'emails' | 'persons'
export type itemSelectedTypeI = 'email' | 'person'

export interface Email {
  id: string
  message_id: string
  date: Date
  from: string
  to: string
  toArray: string[];
  subject: string
  cc: string
  ccArray: string[]
  bcc: string
  x_folder: string
  x_origin: string
  x_file_name: string
  content: string
}

export interface SumaryEmail {
  id: string;
  date: Date;
  day: string;
  time: string;
  from: string;
  to: string;
  toArray: string[];
  subject: string;
}

export interface CardEmailI {
  id: string;
  date: Date;
  day: string;
  time: string;
  from: string;
  to: string;
  subject: string;
}

