export const useHighlight = () => {

  const highlightText = (text: string, term: string): string => {
    if (!term) return text;
    const regex = new RegExp(`(${term})`, 'gi');
    return text.replace(regex, '<span class="highlight">$1</span>');
  };

  const removeHighlightTags = (email: string): string => {
    return email.replace(/<span class="highlight">|<\/span>/g, "");
  };

  return { highlightText, removeHighlightTags };
};
