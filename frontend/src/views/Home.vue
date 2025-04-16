<template>
    <div class="pt-[10%] justify-self-center">
        <h1 class=" text-7xl h-[6rem] bg-clip-text text-transparent bg-gradient-to-r from-blue-400 to-pink-500 tracking-light">
            {{ splashMsg }}
        </h1>
    </div>
</template>

<script setup>
/**
 * IMPORTS
*/
 import { onMounted } from 'vue';
 import router from '../router';


/**
 * PROPS
*/
 const props = defineProps({
    splashMsg: {
        type: String,
        default: "Login to Ledgerly",
    }
 })


/**
 * LIFE CYCLES
*/
 onMounted(() => {
    setInterval(() => {
        // After 4 seconds, navigate to the login page
        router.push("/login");
    }, 4000);
 });

</script>

<style lang="scss" scoped>

</style>