import { CheckIcon, CopyIcon } from "lucide-react";
import { useState } from "react";
import { useNavigate } from "react-router";
import { useApp } from "../App";

export default function URLCard({ url, setUrl }) {
  const navigate = useNavigate();
  const [copied, setCopied] = useState(false);

  const handleCopy = async (data) => {
    try {
      await navigator.clipboard.writeText(data);
      setCopied(true);
      setTimeout(() => setCopied(false), 2000); // reset
    } catch (err) {
      console.error("Failed to copy: ", err);
    }
  };

  return (
    <div className="flex flex-wrap justify-between rounded-2xl shadow shadow-gray-700 p-4">
      <div className="lg:w-1/4 w-full">
        <h4 className="text-gray-400">Original URL</h4>
        {/* <p className="truncate">{url.origin_url}</p> */}
        <a href={url.origin_url}>
          <p className="truncate">{url.origin_url}</p>
        </a>
      </div>
      <div className="lg:w-1/4 w-full">
        <h4 className="text-gray-400">Shorten URL</h4>
        {/* <p>{`${origin}/${url.short_code}`}</p> */}
        <div className="flex gap-2">
          <a href={`${origin}/${url.short_code}`}>
            {`${origin}/${url.short_code}`}{" "}
          </a>
          <span
            onClick={(e) => {
              e.stopPropagation();
              handleCopy(`${origin}/${url.short_code}`);
            }}
          >
            {copied ? (
              <CheckIcon className="h-5 w-5 text-green-600" />
            ) : (
              <CopyIcon className="h-5 w-5 text-gray-400" />
            )}
          </span>
        </div>
      </div>
      <div className="lg:w-1/4 w-1/2">
        <h4 className="text-gray-400">Created At</h4>
        <p>{url.created_at}</p>
      </div>
      <button
        onClick={() => {
          setUrl(url);
          navigate(`/${url.short_code}/details`);
        }}
        className={`px-6 py-3 bg-blue-800 text-white font-medium rounded-md hover:bg-blue-900 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 flex items-center justify-center transition-colors
                  }`}
      >
        Check details
      </button>
    </div>
  );
}
