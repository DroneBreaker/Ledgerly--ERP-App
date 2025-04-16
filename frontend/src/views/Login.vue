<template>
    <div class="justify-items-center pt-[10%]">
        <h1 class="text-7xl bg-clip-text text-transparent bg-gradient-to-r from-blue-500 to-pink-500 tracking-light h-[10rem]">
            {{ loginMsg }}
        </h1>

        <form action="#">
            <input class="rounded-md p-2 px-4 border border-white w-[100%] mb-8 h-[3rem]" type="text" placeholder="Email or Username" id="emailOrUsername"/>

            <input class="rounded-md p-2 px-4 border border-white w-[100%] mb-8 h-[3rem]" type="password" placeholder="Password" id="password"/>
        </form>

        <Button buttonText="Log in"/>
    </div>
</template>

<script setup>
/**
 * IMPORTS
*/
 import Button from '../components/Button.vue';

/**
 * PROPS
 */
 const props = defineProps({
    loginMsg: {
        type: String,
        default: "Login to Ledgerly",
    }
 })


/**
 * FUNCTIONS
*/
 const handleSubmit = () => {
    return;
 }
</script>

<style lang="scss" scoped>

</style>