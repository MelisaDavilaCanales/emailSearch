export type searchTypeI = 'emails' | 'persons'
export type itemSelectedTypeI = 'email' | 'person'

export interface Email {
  id: string
  message_id: string
  date: string
  from: string
  subject: string
  to: string
  cc: string
  toArray: string[]
  ccArray: string[]
  bcc: string
  x_folder: string
  x_origin: string
  x_file_name: string
  content: string
}

export interface SumaryEmail {
  id: string;
  date: string;
  from: string;
  to: string;
  toArray: string[];
  subject: string;
}

export interface CardEmailI {
  id: string;
  date: string;
  from: string;
  to: string;
  subject: string;
}

