<script setup>
// import { UserCircleIcon } from "@heroicons/vue/24/solid";
import { PlusIcon } from "@heroicons/vue/24/outline";
import { RouterLink } from "vue-router";
import SnapcartLogo from "../icons/Snapcart-Logo.vue";
</script>

<template>
    <nav class="bg-gray-900 text-sky-600 flex items-center justify-between px-4 border-b-[1px] border-b-gray-800 sticky top-0 z-50">
        <RouterLink to="/">
            <SnapcartLogo/>
        </RouterLink>
        <div class="flex gap-4">
            <RouterLink class="btn-primary font-semibold flex items-center gap-2" to="/create">
                <PlusIcon class="w-4 h-4"/>
                <span>New Product</span>
            </RouterLink>
            <RouterLink class="btn-primary px-4 py-2" to="/login">
                Login
            </RouterLink>
        </div>
        <!-- <UserCircleIcon class="text-slate-50 h-8 w-8" /> -->
    </nav>
</template>

<style scoped>
    nav{
        height: 64px;
    }
</style>