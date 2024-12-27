export const useCleanTerm = () => {

  const cleanEmail = (email: string): string => {
    return email.replace(/<span class="highlight">|<\/span>/g, "");
  };

  return { cleanEmail };
};
