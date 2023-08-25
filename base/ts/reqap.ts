import to from 'await-to-js';

export interface GrpcError {
  code: number;
  error: string;
  message: string;
  status: number;
  details: any[];
}

export async function reqap<T, U = GrpcError>(
  run: Promise<T>
): Promise<[T | null, U | null]> {
  const [err, res] = await to<T | void | undefined, U>(
    run
    .catch((e) => {
      const res = e.response;

      const invalidStatuses = [302, 401, 403];

      if (invalidStatuses.includes(res.status) || res.type === 'opaqueredirect') {
        window.location.reload();
        return;
      }

      throw res.json();
    })
  );

  if (err) {
    const finalErr: any = await err;
    // Loop over details and check if there is a detail with type
    // `type.googleapis.com/google.rpc.LocalizedMessage`.
    // If locale also matches, then override root message.
    if (finalErr.details) {
      for (const detail of finalErr.details) {
        console.log(detail);
        if (detail['@type'] === 'type.googleapis.com/google.rpc.LocalizedMessage') {
          if (detail.locale === 'en-US') {
            finalErr.message = detail.message;
            break;
          }
        }
      }
    }
    return [null, finalErr];
  }

if (res) {
    return [res, null];
}

  return [null, new Error('Unknown error') as any];
}
