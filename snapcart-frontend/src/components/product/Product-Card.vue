<script setup>
    import { StarIcon, ArrowRightIcon } from "@heroicons/vue/24/solid";
    const props = defineProps(["title", "toggle"])
</script>

<template>
    <div class="bg-slate-900 rounded-md text-slate-50 overflow-hidden">
        <img src="https://placehold.co/600x400" alt="product-image" class="aspect-[3/2]">
        <section class="p-4 flex flex-col">
            <h5 class="text-slate-300">{{ props.title }}</h5>
            <div class="flex items-center justify-between mb-2">
                <h4 class="font-medium">Rp 30.000,00</h4>
                <span class="flex items-center gap-1 text-amber-400">
                    <StarIcon class="h-6 w-6" />
                    <span>4.2/5</span>
                </span>
            </div>
            <p class="line-clamp-2 mb-8 text-slate-400">
                Lorem ipsum dolor sit amet consectetur adipiscing elit. Donec maximus dapibus velit, ac pharetra urna congue at. Duis sed suscipit lacus. Nam aliquet semper malesuada. Fusce gravida nulla non tincidunt tristique. Maecenas eu purus sodales, cursus dolor id, accumsan arcu. Vivamus molestie arcu ut ex euismod convallis. Praesent at lectus tortor. Vivamus tortor lacus, varius id semper at, dapibus vel nisi. Integer rhoncus ligula et dolor gravida, vitae pretium nisi venenatis.
            </p>
            <button class="bg-slate-800 rounded-lg p-2 border-b-2 border-b-slate-950 border-t-2 border-t-slate-700 hover:border-t-slate-600 transition-all active:border-t-slate-950 active:border-b-slate-700" v-on:click="props.toggle">
                <span>See Details <ArrowRightIcon class="h-4 w-4 inline"/></span>
            </button>
        </section>
    </div>
</template>