import { NavLink, Outlet } from 'react-router'

export function Layout() {
  return (
    <>
      <header className="m-auto flex w-7xl justify-between bg-amber-700 p-4 font-mono">
        <h1 className="text-2xl font-extrabold text-white">Stockflow</h1>
        <NavLink to="/" className="text-xl text-white">
          Products
        </NavLink>
      </header>
      <main className="m-auto w-7xl font-mono">
        <Outlet />
      </main>
    </>
  )
}
