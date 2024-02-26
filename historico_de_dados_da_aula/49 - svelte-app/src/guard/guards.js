import { API_HOST } from '../config';

export async function isAuthenticated() {
    const userToken = localStorage.getItem('userToken');

    if(!userToken) return false;
    
    const resposta = await fetch(`${API_HOST}/login/validar`, {
        method: 'HEAD',
        headers: {
            'Authorization': `Bearer ${userToken}`,
            'Content-Type': 'application/json'
        }
    });

    return resposta.ok
}
