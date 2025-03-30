<script lang="ts">
    import { getUserDetails } from "$lib/users";
    import type { PageProps } from "./$types";

    let { data }: PageProps = $props();
    const userId = data.user_id;
    const user = getUserDetails(userId);

    console.log(user);

    const deleteUser = async (id: number) => {
        if (confirm("本当に削除しますか？")) {
            await fetch(`http://localhost:3000/admin/user/${id}`, {
                method: 'DELETE',
                headers: {
                    'Content-Type': 'application/json',
                },
                credentials: "include"
            }).then(res => {
                if (res.status === 200) {
                    location.reload();
                } else {
                    alert("削除に失敗しました");
                }
            }).catch(console.error);
        }
    };
</script>
<h1 class="text-2xl">ユーザー管理</h1>

<table class="table table-zebra">
    <thead>
    <tr>
        <th>ID</th>
        <th>ユーザー名</th>
        <th>ユーザー表示名</th>
        <th>作成日時</th>
        <th>更新日時</th>
        <th>ステータス</th>
        <th>操作</th>
    </tr>
    </thead>
    <tbody>
    {#await user then userData}
        {#if userData !== null}
            <tr>
                <td>{userData.id}</td>
                <td>{userData.name}</td>
                <td>{userData.display_name}</td>
                <td>{userData.createdAt}</td>
                <td>{userData.updatedAt}</td>
                <td>
                    {userData.deletedAt}
                    {#if userData.deletedAt}
                        削除済み
                    {:else}
                        有効
                    {/if}
                </td>
                <td>
                    {#if userData.deletedAt}
                        <button class="btn btn-disabled">削除</button>
                    {:else}
                        <button class="btn btn-error" on:click="{() => {deleteUser(userData.id)}}">削除</button>
                    {/if}
                </td>
            </tr>
        {/if}
        {/await}
    </tbody>
</table>
