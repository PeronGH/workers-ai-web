import { FC } from "hono/jsx";

export const ImageGeneration: FC = () => (
  <div
    x-data="{ prompt: '', imageUrl: null, loading: false }"
    class="container mx-auto p-4"
  >
    <form
      x-data="{async generate(){
      if(!prompt||loading)return;loading=true;
      imageUrl=await fetch('/api/run/@cf/stabilityai/stable-diffusion-xl-base-1.0',
      {method:'POST',body:JSON.stringify({prompt})})
      .then(r=>r.blob()).then(b=>URL.createObjectURL(b)).catch(()=>{}).finally(()=>loading=false);}}"
      class="flex items-center mb-4"
      x-on:submit="event.preventDefault();await generate()"
    >
      <input
        type="text"
        x-model="prompt"
        class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
        placeholder="Type prompt here..."
      />
      <button
        x-bind:disabled="loading"
        class="ml-4 bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
        type="submit"
      >
        <template x-if="loading">
          <span>Loading...</span>
        </template>
        <template x-if="!loading">
          <span>Generate</span>
        </template>
      </button>
    </form>

    <div x-show="imageUrl" class="flex justify-center mt-4">
      <img
        x-bind:src="imageUrl"
        alt="Generated Image"
        class="max-w-full h-auto rounded-md shadow-lg"
      />
    </div>
  </div>
);
