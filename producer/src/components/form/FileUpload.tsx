import { ChangeEventHandler, DragEventHandler, useRef, useState } from "react"
import { parseBlob } from 'music-metadata-browser';
import type { IAudioMetadata } from 'music-metadata-browser'

export default function FileUpload() {

  const [file, setFile] = useState<File>()
  const [metadata, setMetadata] = useState<IAudioMetadata>()
  const [error, setError] = useState<Error>()

  const onChange: ChangeEventHandler<HTMLInputElement> = async (event) => {
    const el = event.target
    if (!el.files) {
      return
    }
    const file = el.files[0]
    setFile(file)
    debugger;
    try {
      const m = await parseBlob(file, {})
      setMetadata(m)
      console.log(m)
    } catch (err) {
      console.log(err)
      setError(new Error("failed to parse mp3 file"))
    }
  }

  const input = useRef<HTMLInputElement>(null)

  const onDrop: DragEventHandler<HTMLDivElement> = (event) => {
    if (!input.current) {
      return
    }
    input.current.files = event.dataTransfer.files;



    event.preventDefault()
  }

  const onDrag: DragEventHandler = (event) => {
    event.preventDefault()
  }

  return <div>
    <label className="block text-sm font-medium text-zinc-300">Track</label>
    <div onDragOver={onDrag} onDragEnter={onDrag} onDrop={onDrop} className="mt-1 flex justify-center px-6 pt-5 pb-6 border-2 border-zinc-300 border-dashed rounded-md">
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
    </div>
  </div>
}