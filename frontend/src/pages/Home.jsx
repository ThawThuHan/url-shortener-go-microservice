import { URLShortenerForm } from "../components/URLShortenForm";
import URLShortenList from "../components/URLShortenList";

export default function Home() {
  return (
    <div className="flex flex-col min-h-screen">
      <div className="container mx-auto px-4 py-4">
        <div>
          <h1 className="text-center text-6xl font-bold">URL Shortener</h1>
        </div>
        <URLShortenerForm />
        <URLShortenList />
      </div>
    </div>
  );
}
