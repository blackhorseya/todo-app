const apiUrl = process.env.REACT_APP_API_URL || '';

function handleResponse(resp) {
  return resp.text().then((body) => {
    const data = body && JSON.parse(body);
    if (400 <= resp.status) {
      const error = (data && data.message) || resp.statusText;
      return Promise.reject(error);
    }

    return data;
  });
}

function list() {
  const reqOpt = {
    method: 'GET',
  };

  return fetch(`${apiUrl}api/v1/tasks?size=10`,
      reqOpt).then(handleResponse);
}

function add(title) {
  const reqOpt = {
    method: 'POST',
    body: JSON.stringify({title}),
  };

  return fetch(`${apiUrl}api/v1/tasks`, reqOpt).then(handleResponse);
}

export const taskServices = {
  list,
  add,
};
