<script setup>
// import { UserCircleIcon } from "@heroicons/vue/24/solid";
import { RouterLink } from "vue-router";
import SnapcartLogo from "../icons/Snapcart-Logo.vue";
</script>

<template>
    <nav class="bg-gray-900 text-sky-600 flex items-center justify-between px-4 border-b-[1px] border-b-gray-800 sticky top-0 z-50">
        <RouterLink to="/">
            <SnapcartLogo/>
        </RouterLink>
        <RouterLink class="btn-primary px-4 py-2" to="/login">
            Login
        </RouterLink>
        <!-- <UserCircleIcon class="text-slate-50 h-8 w-8" /> -->
    </nav>
</template>

<style scoped>
    nav{
        height: 64px;
    }
</style>