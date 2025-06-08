export const usePostComment = () => {
  const postComment = async (articleId: string, comment: { author: string; text: string; created_at: string }) => {
    try {
      const res = await fetch(`http://localhost:8080/articles/${articleId}/comments`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(comment),
      });
      if (!res.ok) throw new Error('Error al enviar el comentario');
      return await res.json();
    } catch (e) {
      throw e;
    }
  };
  return { postComment };
};