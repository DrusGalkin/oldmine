// auth.js - рендерер процесс
import axios from "axios";

const sessName = "sess_id";
const block = document.querySelector('#auth');

// Вспомогательные функции для работы с куками через IPC
async function getCookie(name) {
    if (!window.electronAPI) {
        console.error('Electron API не доступен');
        return null;
    }
    try {
        // Получаем куку из Electron сессии
        const cookie = await window.electronAPI.getCookie(name);
        return cookie ? cookie.value : null;
    } catch (error) {
        console.error('Ошибка при получении куки:', error);
        return null;
    }
}

async function setCookieInElectron(name, value, options = {}) {
    if (!window.electronAPI) {
        console.error('Electron API не доступен');
        return false;
    }

    const cookieData = {
        url: options.url || 'http://localhost',
        name: name,
        value: value,
        domain: options.domain || 'localhost',
        path: options.path || '/',
        secure: options.secure || false,
        httpOnly: options.httpOnly || false,
        expirationDate: options.expires ? Math.floor(Date.now() / 1000) + options.expires : undefined
    };

    try {
        const result = await window.electronAPI.setCookie(cookieData);
        return result.success;
    } catch (error) {
        console.error('Ошибка при установке куки:', error);
        return false;
    }
}

async function checkAuth() {
    if (!window.electronAPI) {
        console.error('Electron API не доступен');
        Login();
        return;
    }

    try {
        // Получаем куку из Electron сессии
        const cookieValue = await getCookie(sessName);
        console.log('Кука из Electron:', cookieValue);

        if (!cookieValue || cookieValue === 'undefined' || cookieValue === 'null') {
            Login();
        } else {
            // Проверяем валидность сессии на сервере
            const response = await axios.post('http://localhost:4000/api/auth/verify', {
                session_id: cookieValue
            });

            if (response.data.valid) {
                Profile(response.data.user);
            } else {
                Login();
            }
        }
    } catch(err) {
        console.error('Ошибка при проверке аутентификации:', err);
        Login();
    }
}

function Profile(user) {
    if (block) {
        block.innerHTML = `
            <h1>Профиль пользователя</h1>
            <div class="profile-info">
                <p><strong>Имя:</strong> ${user.Name || 'Не указано'}</p>
                <p><strong>Email:</strong> ${user.Email}</p>
                <p><strong>ID:</strong> ${user.ID}</p>
                <p><strong>Статус:</strong> ${user.Admin ? 'Администратор' : 'Пользователь'}</p>
                <button id="logoutBtn" class="logout-btn">Выйти</button>
            </div>
        `;

        document.getElementById('logoutBtn').addEventListener('click', handleLogout);
    }
}

function Login() {
    if (block) {
        block.innerHTML = `
            `;

        const form = document.getElementById('loginForm');
        if (form) {
            form.addEventListener('submit', handleLogin);
        }
    } else {
        console.error('Элемент с id="auth" не найден!');
    }
}

async function handleLogin(e) {
    e.preventDefault();
    const email = document.getElementById('email').value;
    const password = document.getElementById('password').value;

    try {
        const response = await axios.post('http://localhost:4000/api/auth/login', {
            email: email,
            password: password
        });

        // Сохраняем куку в Electron сессии
        await setCookieInElectron(sessName, response.data.session_id, {
            expires: 7 * 24 * 60 * 60 // 7 дней в секундах
        });

        // Также сохраняем в localStorage для совместимости
        localStorage.setItem('admin', response.data.user.Admin);
        localStorage.setItem('email', response.data.user.Email);
        localStorage.setItem('id', response.data.user.ID);
        localStorage.setItem('name', response.data.user.Name);
        localStorage.setItem('payment', response.data.user.Payment);

        checkAuth();
    } catch (error) {
        console.error('Ошибка при входе:', error);
        alert('Ошибка входа: ' + (error.response?.data?.message || error.message));
    }
}

async function handleLogout() {
    try {
        // Удаляем куку из Electron сессии
        if (window.electronAPI) {
            await window.electronAPI.removeCookie(sessName);
        }

        // Очищаем localStorage
        localStorage.clear();

        // Возвращаем на страницу логина
        Login();
    } catch (error) {
        console.error('Ошибка при выходе:', error);
    }
}

// Инициализация
document.addEventListener('DOMContentLoaded', () => {
    checkAuth();
});

// Экспортируем функции для тестирования (опционально)
export { checkAuth, getCookie, setCookieInElectron };