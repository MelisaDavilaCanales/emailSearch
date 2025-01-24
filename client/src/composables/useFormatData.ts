export const useFormatData = () => {
  const formatDate = (isoDate: string): string => {
    const [datePart, timePart] = isoDate.split('T')
    const day = datePart
    const time = timePart.split(':').slice(0, 2).join(':')
    const formattedDate = day + ' ' + time
    return formattedDate
  }

  const convertStringListToArray = (emails: string): string[] => {
    return emails.split(',').map((email: string) => email.trim())
  }

  const cleanQuery = (query: string): string => {
    return query.replace(/[{}"":*]/g, '')
  }

  return { formatDate, convertStringListToArray, cleanQuery }
}
