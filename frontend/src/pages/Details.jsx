import { useQuery } from "react-query";
import { Link, useNavigate, useParams } from "react-router";
import { getAccessLog } from "../lib/fetch";
import { ArrowLeftIcon } from "lucide-react";
import { useApp } from "../App";

export default function Details() {
  const { url } = useApp();
  const navigate = useNavigate();
  const { short_code } = useParams();
  const { data, error } = useQuery("access_log", () =>
    getAccessLog(short_code)
  );

  const result = data?.logs;

  const origin = window.location.origin;

  return (
    <div className="container p-4 mx-auto my-4">
      <div>
        <button
          onClick={() => navigate(-1)}
          className="mt-4 inline-flex items-center text-blue-600 hover:text-blue-800"
        >
          <ArrowLeftIcon className="h-4 w-4 mr-2" />
          Back to Home
        </button>
      </div>
      {url == null ? (
        <div className="flex justify-center items-center w-full h-50">
          <h3 className="text-2xl text-center text-red-500">
            Something Wrong! please go back to{" "}
            <Link to={"/"} className="underline">
              Home Page
            </Link>
          </h3>
        </div>
      ) : (
        <div className="flex flex-col rounded-2xl border p-4">
          <h3 className="text-2xl font-bold">URL Detail</h3>
          <div className="pt-4">
            <h3 className="text-xl text-gray-500">Original URL</h3>
            <p>{url.origin_url}</p>
          </div>
          <div className="pt-4">
            <h3 className="text-xl text-gray-500">Shorten URL</h3>
            <p>{`${origin}/${url.short_code}`}</p>
          </div>
          <div className="w-1/2 flex justify-between pt-4">
            <div>
              <h3 className="text-xl text-gray-500">Created At</h3>
              <p>{url.created_at}</p>
            </div>
            <div>
              <h3 className="text-xl text-gray-500">Total Click</h3>
              <p>{result?.length || 0}</p>
            </div>
          </div>
          <h3 className="text-2xl pt-4">Access Logs</h3>
          <div className="overflow-x-auto pt-4">
            <table className="min-w-full divide-y border">
              <thead className="">
                <tr>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-300 uppercase tracking-wider">
                    ID
                  </th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-300 uppercase tracking-wider">
                    Time
                  </th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-300 uppercase tracking-wider">
                    IP Address
                  </th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-300 uppercase tracking-wider">
                    City
                  </th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-300 uppercase tracking-wider">
                    Location (Lat/Long)
                  </th>
                </tr>
              </thead>
              <tbody className="divide-y ">
                {data?.logs?.map((log) => (
                  <tr key={log.id}>
                    <td className="px-6 py-4 whitespace-nowrap text-sm">
                      {log.id}
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm">
                      {log.access_time}
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm">
                      {log.ip_address}
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm truncate max-w-xs">
                      {log.city}
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm truncate max-w-xs">
                      {log.location}
                    </td>
                  </tr>
                ))}
                {result?.length === 0 ||
                  (result?.length === undefined && (
                    <tr>
                      <td
                        colSpan={4}
                        className="px-6 py-4 text-center text-sm text-gray-500"
                      >
                        No access logs yet
                      </td>
                    </tr>
                  ))}
              </tbody>
            </table>
          </div>
        </div>
      )}
    </div>
  );
}
