export type searchTypeI = 'emails' | 'persons'
export type itemSelectedTypeI = 'email' | 'person'

export interface IEmail {
  id: string
  message_id: string
  date: string
  from: string
  subject: string
  to: string
  cc: string
  toArray: string[]
  ccArray: string[]
  x_folder: string
  x_origin: string
  x_file_name: string
  content: string
}

export interface IHighlightedEmailDetail {
  id: string
  message_id: string
  date: string
  from: string
  subject: string
  toArray: string[]
  ccArray: string[]
  x_folder: string
  x_origin: string
  x_file_name: string
  content: string
}

export interface ISumaryEmail {
  id: string
  date: string
  from: string
  to: string
  toArray: string[]
  subject: string
}

export interface ICardEmail {
  id: string
  date: string
  from: string
  to: string
  subject: string
}

export interface IPerson {
  id: string
  name: string
  email: string
}

export interface IRequestError {
  status: boolean
  message: string
}

export interface IServerErrorResponse {
  status: boolean
  code: number
  message: string
}
