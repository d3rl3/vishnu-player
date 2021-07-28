import {registerRoute} from 'workbox-routing';
import {NetworkFirst,CacheFirst,} from 'workbox-strategies'
import { CacheableResponsePlugin } from 'workbox-cacheable-response';

// registerRoute(
//     ({ request }) => request.destination === 'image',
//     new CacheFirst());

// Cache page navigations (html) with a Network First strategy
registerRoute(
    // Check to see if the request is a navigation to a new page
    ({request}) => request.mode === 'navigate',
    // Use a Network First caching strategy
    new NetworkFirst({
      // Put all cached files in a cache named 'pages'
      cacheName: 'pages',
      plugins: [
        // Ensure that only requests that result in a 200 status are cached
        new CacheableResponsePlugin({
          statuses: [200],
        }),
      ],
    }),
);
