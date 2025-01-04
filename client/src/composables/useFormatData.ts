export const useFormatData = () => {
  const formatDate = (isoDate: string): string => {
    const [datePart, timePart] = isoDate.split('T')
    const day = datePart
    const time = timePart.split(':').slice(0, 2).join(':') // Extrae solo hh:mm
    const formattedDate = day + ' ' + time
    return formattedDate
  }

  const convertToArray = (emails: string): string[] => {
    return emails.split(',').map((email: string) => email.trim())
  }

  return { formatDate, convertToArray }
}
