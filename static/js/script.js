document.addEventListener("htmx:afterRequest", function(event) {
  document.querySelectorAll(".sidebar-btn").forEach(btn => {
    btn.classList.add("text-white");
    btn.classList.remove("text-orange-400");
  });

  const btn = event.detail.elt;

  if (btn && btn.classList.contains("sidebar-btn")) {
    btn.classList.remove("text-white");
    btn.classList.add("text-orange-400");
  }
});

