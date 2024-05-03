<script setup>
    import { StarIcon, ArrowRightIcon } from "@heroicons/vue/24/solid";
    const props = defineProps(["data", "toggle", "id"])
</script>

<template>
    <div class="bg-gray-900 rounded-md text-slate-50 overflow-hidden border-[1px] border-gray-900 max-w-[300px]">
        <section class="w-full aspect-square relative">
            <img :src="props.data.img" alt="product-image" class="h-full w-full object-cover">
            <span class="flex items-center gap-1 text-amber-400 absolute top-4 left-4 bg-gray-800 py-1 px-2 rounded-md">
                <StarIcon class="h-4 w-4" />
                <span>{{props.data.rating}}/5</span>
            </span>
            <div class="absolute bottom-0 p-2 z-10">
                <h3 class="font-semibold">Rp {{ props.data.price.toLocaleString() }}</h3>
                <p class="text-gray-300 text-base line-clamp-1">{{ props.data.title }}</p>
            </div>
            <div class="bg-gradient-to-t from-gray-900 via-gray-900/75 via-75%% to-transparent w-full h-1/3 absolute bottom-0"></div>
        </section>
        <section class="p-4 pt-0 flex flex-col">
            <hr class="mb-4 border-gray-800">
            <button class="btn-primary py-2" v-on:click="()=>props.toggle(props.id)">
                <span>See Details <ArrowRightIcon class="h-4 w-4 inline"/></span>
            </button>
        </section>
    </div>
</template>