export const login = async (email: string, password: string) => {
    const res = await fetch('http://localhost:3000/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email: email, password: password }),
        // fetch APIはデフォルトでセッション情報を送信しないため、
        // 明示的に送信するように設定する
        credentials: "include"
    }).catch(console.error);

    if (res) {
        console.log(await res.json());
    } else {
        console.error(res);
    }
}

export const signup = async (name: string, displayName: string, email: string, password: string) => {
    const res = await fetch('http://localhost:3000/signup', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            name: name,
            display_name: displayName,
            email: email,
            password: password
        }),
    }).catch(console.error);

    if (res) {
        console.log(await res.json());
    } else {
        console.error(res);
    }
};
