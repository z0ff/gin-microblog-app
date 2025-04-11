<script lang="ts">
    import { onMount } from 'svelte';
    import { login } from "$lib/auth";

    let email: HTMLInputElement;
    let password: HTMLInputElement;

    let loginButtonEnabled = false;

    onMount(() => {
        email = document.getElementById('form-email') as HTMLInputElement;
        password = document.getElementById('form-password') as HTMLInputElement;
    });

    const onClickLoginButton = async () => {
        const result = await login(email.value.trim(), password.value.trim());

        email.value = "";
        password.value = "";

        if (result === false) {
            alert("login failed");
        } else {
            location.href = "/";
        }
    };

    const onInput = () => {
        toggleEnabledLoginButton();
    }

    const toggleEnabledLoginButton = () => {
        loginButtonEnabled = email.value.trim().length > 0 && password.value.trim().length > 0;
    }
</script>

<div class="w-full flex justify-center items-center">
    <div class="card card-body flex flex-column max-w-sm">
        <input type="text" id="form-email" placeholder="e-mail" on:input={onInput} class="input" />
        <input type="password" id="form-password" placeholder="password" on:input={onInput} class="input" />
        <button type="button" on:click={onClickLoginButton} disabled="{!loginButtonEnabled}" class="btn btn-primary">ログイン</button>
        <a class="btn btn-primary" href="/signup">会員登録</a>
    </div>
</div>