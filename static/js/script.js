document.addEventListener('htmx:afterRequest', function() {
  const currentPath = window.location.pathname;

  fetch('/api/filetree', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ path: currentPath })
  })
    .then(response => {
      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
      return response.json();
    })
    .then(fileTree => {
      console.log('File tree data:', fileTree);
      updateDOM(fileTree);
    })
    .catch(error => {
      console.error('Fetch error:', error);
    });
});

function updateDOM(node) {
  const domNode = document.querySelector(`[data-path="${node.path}"]`);
  if (domNode) {
    if (node.isActive) {
      domNode.classList.add('text-orange-400');
      domNode.classList.remove('text-white');
    } else {
      domNode.classList.add('text-white');
      domNode.classList.remove('text-orange-400');
    }
  }

  if (node.children && node.children.length > 0) {
    node.children.forEach(childNode => {
      updateDOM(childNode);
    });
  }
}

