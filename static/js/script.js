document.addEventListener("htmx:afterRequest", function(event) {
  document.querySelector(".sidebar-btn").forEach(btn => {
    btn.classList.add("text-white");
    btn.classList.remove("text-orange-400");
  })

  const btn = event.detail.elt
  btn.classList.remove("text-white")
  btn.classList.add("text-orange-400")

});
