export const useAuth = () => {
  const login = async (email: string, password: string) => {
    try {
      const res = await fetch('http://localhost:8080/login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, password })
      });
      if (!res.ok) throw new Error('Credenciales incorrectas');
      return await res.json();
    } catch (e) {
      throw e;
    }
  };

  const register = async (name: string, email: string, password: string) => {
    try {
      const res = await fetch('http://localhost:8080/signup', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ name, email, password })
      });
      if (!res.ok) throw new Error('Error al registrar usuario');
      return await res.json();
    } catch (e) {
      throw e;
    }
  };

  return { login, register };
};
