import Footer from "@/components/Footer";
import { fontSans } from "@/lib/fonts";
import "@/styles/globals.css";
import AnnouncementBar from '@/components/AnnouncementBar';
import Navbar from "@/components/Navbar";
import { ThemeProvider } from "@/components/ThemeProvider";
import { cn } from "@/lib/utils"
import type { Props } from "@/types/RootLayout"

export const metadata = {
  title: "Thích Tài Liệu",
  description: "test thu description",
}

export default function RootLayout({ children }: Props): JSX.Element {
  return (
    <html lang="vi">
      <body className={cn("min-h-screen bg-background font-sans antialiased", fontSans.variable)}>
        <ThemeProvider attribute="class" defaultTheme="system" enableSystem>
          <div className="relative min-h-screen">
            <AnnouncementBar />
            <Navbar />
            <div className="container">
              <div className="py-12">{children}</div>
            </div>
            <Footer />
          </div>
        </ThemeProvider>
      </body>
    </html>
  )
}