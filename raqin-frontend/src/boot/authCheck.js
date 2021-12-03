export default async ({ store, router }) => {
    store.dispatch("auth/myProfile")
        .catch((err) => {
            console.log("err is", err);
            router.push("/login");
        });

}