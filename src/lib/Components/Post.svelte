<script>
    import { strToDate } from "$lib/datetime";
    export let post;
    const postID = post.ID;
    let isLiked = post.IsLiked;

    const onLikeBtnClicked = async () => {
        const res = await fetch(`http://localhost:3000/like/${postID}`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            // fetch APIはデフォルトでセッション情報を送信しないため、
            // 明示的に送信するように設定する
            credentials: "include"
        }).catch(console.error);

        if (res) {
            console.log(await res.json());
            isLiked = true;
        } else {
            console.error(res);
        }
    }

    const onUnLikeBtnClicked = async () => {
        const res = await fetch(`http://localhost:3000/unlike/${postID}`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            // fetch APIはデフォルトでセッション情報を送信しないため、
            // 明示的に送信するように設定する
            credentials: "include"
        }).catch(console.error);

        if (res) {
            console.log(await res.json());
            isLiked = false;
        } else {
            console.error(res);
        }
    }
</script>
<div class="card card-bordered bg-base-50 shadow-lg">
    <div class="card-body">
        <a href="/@{post.User.Name}"><span class="font-bold">{post.User.DisplayName}</span> @{post.User.Name}</a>
        <p>{strToDate(post.CreatedAt)}</p>
        <p class="whitespace-pre-wrap">{post.Content}</p>
        <button class="btn {isLiked ? 'btn-primary' : 'btn-ghost'}" on:click={isLiked ? onUnLikeBtnClicked : onLikeBtnClicked}>{isLiked ? "Unlike" : "Like"}</button>
    </div>
</div>
