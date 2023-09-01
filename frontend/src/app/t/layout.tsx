export default function Layout({ children }: { children: React.ReactNode }) {
  return (
    <div className="container">
      <div className="py-6">
        <div>{children}</div>
      </div>
    </div>
  )
}
