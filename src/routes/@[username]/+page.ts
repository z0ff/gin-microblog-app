import type { PageLoad } from '@sveltejs/kit';
import {getMe, getUser} from "$lib/auth";
import type {User} from "$lib/types";

export const load: PageLoad = async ({ params }) => {
    // let user: User | null = null;
    // let isMe = false;
    //
    // const me = await getMe();
    // if (me === null) {
    //     return { error: "user not found" };
    // }
    //
    // console.log(params.username);
    // console.log(me);
    //
    // if (params.username === me.name) {
    //     isMe = true;
    //     user = me;
    // } else {
    //     user = await getUser(params.username);
    //     if (user === null) {
    //         return { error: "user not found" };
    //     }
    // }
    //
    // return {
    //     props: {
    //         user: user,
    //         isMe: isMe
    //     }
    // }

    return {
        props: {
            username: params.username
        }
    }
}
