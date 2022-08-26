import { ButtonHTMLAttributes, ChangeEventHandler, DragEventHandler, MouseEventHandler, useRef, useState } from "react"
import { parseBlob } from 'music-metadata-browser';
import type { IAudioMetadata } from 'music-metadata-browser'

export default function FileUpload() {

  const [file, setFile] = useState<File>()
  const [metadata, setMetadata] = useState<IAudioMetadata>()
  const [error, setError] = useState<Error>()



  const onChange: ChangeEventHandler<HTMLInputElement> = (event) => {
    const el = event.target
    if (!el.files) {
      return
    }
    const file = el.files[0]
    setFile(file)
    parseAudioFile(file)
  }

  const input = useRef<HTMLInputElement>(null)

  const onDrop: DragEventHandler<HTMLDivElement> = (event) => {
    event.preventDefault()
    if (!input.current) {
      return
    }
    input.current.files = event.dataTransfer.files;
    const file = event.dataTransfer.files[0]
    setFile(file)
    parseAudioFile(file)
  }

  const onDrag: DragEventHandler = (event) => {
    event.preventDefault()
  }

  const parseAudioFile = (file: File) => {
    if(!file) {
      return
    }
    parseBlob(file, {}).then((metadata) => setMetadata(metadata)).catch((err) => setError(err))
  }

  if (!file || !metadata) {
    return (<div onDrop={onDrop} onDragOver={onDrag} onDragEnter={onDrag} className="mt-1 flex justify-center px-6 pt-8 pb-10 border-2 border-zinc-300 border-dashed hover:border-zinc-600 hover:shadow-md rounded-md transition-all hover:rounded-xl">
      <div className="space-y-1 text-center">
        <i className='font-icons text-6xl not-italic'>ß</i>
        <div className="flex text-sm text-zinc-400">
          <label
            htmlFor="file-upload"
            className="relative cursor-pointer rounded-md font-medium text-zinc-700 hover:text-zinc-400 focus-within:outline-none focus-within:ring-2 focus-within:ring-offset-2 focus-within:ring-zinc-500"
          >
            <span>Upload a file</span>
            <input ref={input} id="file-upload" name="file-upload" type="file" className="sr-only" onChange={onChange} />
          </label>
          <p className="pl-1">or drag and drop</p>
        </div>
        <p className="text-xs text-zinc-500">MP3 up to 50MB</p>
      </div>
    </div>)
  }

  const cover = metadata?.common?.picture;
  let imageData = ""
  if (cover && cover[0]) {
    imageData = `data:${cover[0].format};base64, ${cover[0].data.toString("base64")}`
  }

  const clearInput: MouseEventHandler<HTMLButtonElement> = (event) => {
    setFile(undefined)
    setMetadata(undefined)
    setError(undefined)
  }

  return <div>
    <div className="mt-1 flex px-6 pt-5 pb-6 border border-zinc-100 rounded-md">
      <div className="">
        <div className="md:grid grid-cols-2 gap-4">
          <img className="max-w-sm" src={imageData} />
          <div className="text-left flex flex-col">
            <div className="mb-2">
              <span className="text-sm block font-medium text-zinc-500">Title</span>
              <span>{metadata.common.title}</span>
            </div>
            <div className="mb-2">
              <span className="text-sm block font-medium text-zinc-500">Artist</span>
              <span>{metadata.common.artist}</span>
            </div>
            <div className="mb-2">
              <span className="text-sm block font-medium text-zinc-500">Album</span>
              <span>{metadata.common.album}</span>
            </div>
            <div className="flex-grow"></div>
            <div>
              <button type="button" onClick={clearInput}>Cancel</button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
}