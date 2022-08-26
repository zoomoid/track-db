import React from 'react';
import logo from './logo.svg';
import './App.css';
import FileUpload from './components/form/FileUpload';

function App() {
  return (
    <div className="flex flex-col grow bg-white text-black">
      <header className="">
        <div className='mx-auto py-1 container flex'>
          <a href='/'>
            <p className='uppercase font-extrabold text-2xl'>Producer</p>
          </a>
          <div className='flex-grow'></div>
          <a href="/add">
            New Track
          </a>
        </div>
      </header>
      <main className="container flex-grow mx-auto">
        <section className='mt-32'>
          <h1 className='uppercase tracking-tighter font-title font-bold text-8xl'>New Track</h1>
          <form>
            <FileUpload></FileUpload>
          </form>
        </section>

      </main>
      <footer className="">

      </footer>
    </div >
  );
}

export default App;
