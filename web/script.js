const listEl = document.querySelector("#slangList");
const resultCountEl = document.querySelector("#resultCount");
const searchInput = document.querySelector("#searchInput");
const sceneSelect = document.querySelector("#sceneSelect");
const emotionSelect = document.querySelector("#emotionSelect");
const randomButton = document.querySelector("#randomButton");
const randomSection = document.querySelector("#randomSection");
const randomSlangEl = document.querySelector("#randomSlang");
const closeRandomButton = document.querySelector("#closeRandomButton");
const slangForm = document.querySelector("#slangForm");
const formMessage = document.querySelector("#formMessage");
const loadMoreWrap = document.querySelector("#loadMore");
const loadMoreButton = document.querySelector("#loadMoreButton");

const PAGE_SIZE = 20;
let currentOffset = 0;

async function fetchJSON(url, options) {
  const res = await fetch(url, options);
  const data = await res.json();
  if (!res.ok) throw new Error(data.error || "API request failed");
  return data;
}

function buildSlangCard(slang) {
  const article = document.createElement("article");
  article.className = "slangCard";

  const title = document.createElement("h3");
  title.textContent = slang.slang;
  article.appendChild(title);

  for (const m of slang.meanings || []) {
    const meaningEl = document.createElement("div");
    meaningEl.className = "meaning";

    const meaningJa = document.createElement("p");
    meaningJa.className = "meaningJa";
    meaningJa.textContent = m.meaning_ja;
    meaningEl.appendChild(meaningJa);

    if (m.nuance_ja) {
      const nuance = document.createElement("p");
      nuance.className = "nuance";
      nuance.textContent = `ニュアンス: ${m.nuance_ja}`;
      meaningEl.appendChild(nuance);
    }

    for (const ex of m.examples || []) {
      if (ex.sentence_en) {
        const p = document.createElement("p");
        p.className = "example";
        p.textContent = `Example: ${ex.sentence_en}`;
        meaningEl.appendChild(p);
      }
      if (ex.sentence_ja) {
        const p = document.createElement("p");
        p.className = "example";
        p.textContent = `訳: ${ex.sentence_ja}`;
        meaningEl.appendChild(p);
      }
    }

    const scenes = m.scene || [];
    const emotions = m.emotion_categories || [];
    if (scenes.length > 0 || emotions.length > 0) {
      const meta = document.createElement("div");
      meta.className = "meta";
      scenes.forEach((tag) => {
        const pill = document.createElement("span");
        pill.className = "pill pill-scene";
        pill.textContent = tag;
        meta.appendChild(pill);

        pill.addEventListener("click", () => {
          sceneSelect.value = tag;
          searchInput.value = "";
          refreshList();
        });
      });
      emotions.forEach((tag) => {
        const pill = document.createElement("span");
        pill.className = "pill pill-emotion";
        pill.textContent = tag;
        meta.appendChild(pill);

        pill.addEventListener("click", () => {
          emotionSelect.value = tag;
          searchInput.value = "";
          refreshList();
        });
      });
      meaningEl.appendChild(meta);
    }

    if (m.warning_ja) {
      const warning = document.createElement("div");
      warning.className = "warning";
      warning.textContent = `注意: ${m.warning_ja}`;
      meaningEl.appendChild(warning);
    }

    article.appendChild(meaningEl);
  }

  const relatedBtn = document.createElement("button");
  relatedBtn.type = "button";
  relatedBtn.className = "relatedButton";
  relatedBtn.textContent = "関連スラングを見る";

  const relatedArea = document.createElement("div");
  relatedArea.className = "relatedArea hidden";

  let relatedLoaded = false;

  relatedBtn.addEventListener("click", () => {
    const isOpen = !relatedArea.classList.contains("hidden");
    if (isOpen) {
      // 閉じる
      relatedArea.classList.add("hidden");
      relatedBtn.textContent = "関連スラングを見る";
    } else {
      // 開く（初回のみ fetch）
      relatedArea.classList.remove("hidden");
      relatedBtn.textContent = "関連スラングを閉じる";
      if (!relatedLoaded) {
        relatedLoaded = true;
        loadRelatedSlangs(slang.id, relatedArea);
      }
    }
  });

  article.appendChild(relatedBtn);
  article.appendChild(relatedArea);
  return article;
}

async function loadSlangs(append = false) {
  if (!append) currentOffset = 0;

  try {
    const data = await fetchJSON(
      `/api/slangs?limit=${PAGE_SIZE}&offset=${currentOffset}`,
    );

    if (!append) {
      listEl.innerHTML = "";
      resultCountEl.textContent = `${data.length} 件`;
    } else {
      const prev = parseInt(resultCountEl.textContent) || 0;
      resultCountEl.textContent = `${prev + data.length} 件`;
    }

    if (data.length === 0 && !append) {
      listEl.innerHTML = `<div class="empty">スラングがありません。</div>`;
      loadMoreWrap.classList.add("hidden");
      return;
    }

    data.forEach((s) => listEl.appendChild(buildSlangCard(s)));
    currentOffset += data.length;
    loadMoreWrap.classList.toggle("hidden", data.length < PAGE_SIZE);
  } catch (err) {
    listEl.innerHTML = `<div class="error">${err.message}</div>`;
  }
}

async function searchSlangs() {
  const keyword = searchInput.value.trim();
  try {
    const data = await fetchJSON(
      `/api/slangs/search?keyword=${encodeURIComponent(keyword)}`,
    );
    listEl.innerHTML = "";
    resultCountEl.textContent = `${data.count} 件`;
    loadMoreWrap.classList.add("hidden");
    if (data.count === 0) {
      listEl.innerHTML = `<div class="empty">該当するスラングがありません。</div>`;
      return;
    }
    data.items.forEach((s) => listEl.appendChild(buildSlangCard(s)));
  } catch (err) {
    listEl.innerHTML = `<div class="error">${err.message}</div>`;
  }
}

async function loadCategories() {
  try {
    const scenes = await fetchJSON("/api/categories/scenes");
    scenes.forEach((s) => {
      const opt = document.createElement("option");
      opt.value = s;
      opt.textContent = s;
      sceneSelect.appendChild(opt);
    });
  } catch (_) {}

  try {
    const emotions = await fetchJSON("/api/categories/emotion_categories");
    emotions.forEach((e) => {
      const opt = document.createElement("option");
      opt.value = e;
      opt.textContent = e;
      emotionSelect.appendChild(opt);
    });
  } catch (_) {}
}

async function filterSlangs() {
  const scene = sceneSelect.value;
  const emotion = emotionSelect.value;
  const params = new URLSearchParams();
  if (scene) params.set("scene", scene);
  if (emotion) params.set("emotion", emotion);

  try {
    const data = await fetchJSON(`/api/slangs/filter?${params.toString()}`);
    listEl.innerHTML = "";
    loadMoreWrap.classList.add("hidden");
    const items = data || [];
    resultCountEl.textContent = `${items.length} 件`;
    if (items.length === 0) {
      listEl.innerHTML = `<div class="empty">該当するスラングがありません。</div>`;
      return;
    }
    items.forEach((s) => listEl.appendChild(buildSlangCard(s)));
  } catch (err) {
    listEl.innerHTML = `<div class="error">${err.message}</div>`;
  }
}

async function loadRelatedSlangs(slangId, container) {
  container.innerHTML = `<p class="nuance">読み込み中...</p>`;
  try {
    const data = await fetchJSON(`/api/slangs/${slangId}/related`);
    container.innerHTML = "";
    const items = data.items || data || [];   // レスポンス形式に合わせて調整
    if (items.length === 0) {
      container.innerHTML = `<p class="nuance">関連スラングはありません。</p>`;
      return;
    }
    items.forEach((s) => container.appendChild(buildSlangCard(s)));
  } catch (err) {
    container.innerHTML = `<div class="error">${err.message}</div>`;
  }
}

function refreshList() {
  const keyword = searchInput.value.trim();
  const scene = sceneSelect.value;
  const emotion = emotionSelect.value;

  if (keyword) {
    searchSlangs();
  } else if (scene || emotion) {
    filterSlangs();
  } else {
    loadSlangs();
  }
}

async function showRandomSlang() {
  try {
    const data = await fetchJSON("/api/slangs/random");
    randomSlangEl.innerHTML = "";
    if (data.items && data.items.length > 0) {
      randomSlangEl.appendChild(buildSlangCard(data.items[0]));
    }
    randomSection.classList.remove("hidden");
  } catch (err) {
    randomSlangEl.innerHTML = `<div class="error">${err.message}</div>`;
    randomSection.classList.remove("hidden");
  }
}

function debounce(fn, delay) {
  let id;
  return () => {
    clearTimeout(id);
    id = setTimeout(fn, delay);
  };
}

slangForm.addEventListener("submit", async (e) => {
  e.preventDefault();
  formMessage.textContent = "追加中...";

  const fd = new FormData(slangForm);
  const sceneTags = (fd.get("scene") || "")
    .split(",")
    .map((s) => s.trim())
    .filter(Boolean);
  const emotionTags = (fd.get("emotion") || "")
    .split(",")
    .map((s) => s.trim())
    .filter(Boolean);
  const sentenceEn = (fd.get("sentence_en") || "").trim();
  const sentenceJa = (fd.get("sentence_ja") || "").trim();
  const examples =
    sentenceEn || sentenceJa
      ? [{ sentence_en: sentenceEn, sentence_ja: sentenceJa }]
      : [];

  const payload = {
    id: fd.get("id").trim(),
    slang: fd.get("slang").trim(),
    meanings: [
      {
        meaning_ja: fd.get("meaning_ja").trim(),
        nuance_ja: (fd.get("nuance_ja") || "").trim(),
        examples,
        scene: sceneTags,
        emotion_categories: emotionTags,
        warning_ja: (fd.get("warning_ja") || "").trim(),
      },
    ],
  };

  try {
    await fetchJSON("/api/slangs/add", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(payload),
    });
    slangForm.reset();
    formMessage.textContent = "スラングを追加しました。";
    await loadSlangs();
  } catch (err) {
    formMessage.textContent = err.message;
  }
});

const debouncedRefresh = debounce(refreshList, 300);
searchInput.addEventListener("input", debouncedRefresh);
sceneSelect.addEventListener("change", refreshList);
emotionSelect.addEventListener("change", refreshList);
randomButton.addEventListener("click", showRandomSlang);
closeRandomButton.addEventListener("click", () =>
  randomSection.classList.add("hidden"),
);
loadMoreButton.addEventListener("click", () => loadSlangs(true));

document.querySelector("h1").addEventListener("click", () => {
  searchInput.value = "";
  sceneSelect.value = "";
  emotionSelect.value = "";
  window.scrollTo({ top: 0, behavior: "smooth" });
  loadSlangs();
});

loadCategories();
loadSlangs();