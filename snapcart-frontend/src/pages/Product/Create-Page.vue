<script setup>
    import { ArrowLeftIcon } from "@heroicons/vue/24/solid";
    import { PlusIcon, XMarkIcon } from "@heroicons/vue/24/outline";
    import { ref, watch} from "vue";

    const file = ref({url:"", file:null})
    const fileHandler = e => {
        if(e.target.files){
            const newFile = e.target.files[0];
            file.value = {url : URL.createObjectURL(newFile), file : newFile};
        }
    }

    watch(file, () => {
        console.log(file.value);
    })
</script>

<template>
    <form class="min-h-screen">
        <section class="p-4 border-b-[1px] border-b-gray-800 flex justify-between items-center">
            <RouterLink class="flex hover:gap-3 gap-2 text-white items-center font-semibold transition-all" to="/">
                <ArrowLeftIcon class="w-4 h-4"/>
                <span>Back</span>
            </RouterLink>
            <div class="flex gap-4 items-center border-l-[1px] border-gray-800 pl-4">
                <button class="btn-secondary flex gap-2 items-center" type="button">
                    <XMarkIcon class="w-4 h-4"/>
                    <span>Cancel</span>
                </button>
                <button class="btn-primary flex gap-2 items-center">
                    <PlusIcon class="w-4 h-4"/>
                    <span>Add Product</span>
                </button>
            </div>
        </section>
        <section class="p-4">
            <h2 class="font-bold mb-4 text-white">New Product</h2>
            <div class="flex gap-4">
                <main class="w-3/5 flex flex-col gap-4">
                    <div>
                      <label for="title" class="block text-sm font-medium leading-6 text-gray-200">Product Title</label>
                      <div class="mt-2">
                        <input 
                            id="title" 
                            name="title" 
                            type="text" 
                            required 
                            class="block w-full rounded-md border-0 py-1 px-1.5 placeholder:text-gray-400 sm:text-sm sm:leading-6 bg-gray-800 ring-1 ring-gray-700 text-white"
                            placeholder="Enter new product title...">
                      </div>
                    </div>
                    <div>
                      <label for="price" class="block text-sm font-medium leading-6 text-gray-200">Price</label>
                      <div class="mt-2 flex">
                        <span class="bg-gray-800 ring-1 ring-gray-700 text-white py-1.5 px-2 rounded-l-lg">Rp. </span>
                        <input id="price" name="price" type="number" min="0" placeholder="Enter product price..." required class="block w-full rounded-md border-0 py-1 px-1.5 placeholder:text-gray-400 sm:text-sm sm:leading-6 bg-gray-800 ring-1 ring-gray-700 text-white rounded-l-none">
                      </div>
                    </div>
                    <div>
                      <label for="description" class="block text-sm font-medium leading-6 text-gray-200">Description</label>
                      <div class="mt-2">
                        <textarea id="description" name="description" required placeholder="Write product description..." class="block w-full rounded-md border-0 py-1 px-1.5 placeholder:text-gray-400 sm:text-sm sm:leading-6 bg-gray-800 ring-1 ring-gray-700 text-white min-h-64"/>
                      </div>
                    </div>
                </main>
                <aside class="w-2/5 border-l-[1px] border-l-gray-800 pl-4">
                    <div v-if="file.file == null" class="w-full aspect-video flex items-center justify-center border-2 border-dashed rounded-md text-gray-500 mb-4 border-gray-500">
                        Preview Image
                    </div>
                    <img v-else :src="file.url" alt="filePreview" class="w-full aspect-video object-contain border-[1px] border-dashed rounded-md mb-4">
                    <div>
                      <label for="image" class="block text-sm font-medium leading-6 text-gray-200">Image</label>
                      <div class="mt-2">
                        <input id="image" name="image" type="file" @change="fileHandler" required class="block w-full rounded-md border-0 py-1 px-1.5 placeholder:text-gray-400 sm:text-sm sm:leading-6 bg-gray-800 ring-1 ring-gray-700 text-white">
                      </div>
                    </div>
                </aside>
            </div>
        </section>
    </form>
</template>