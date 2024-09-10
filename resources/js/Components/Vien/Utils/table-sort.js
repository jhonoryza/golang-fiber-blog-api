import { router } from "@inertiajs/vue3";
import { ref } from "vue";
import {mergeURLSearchParams, objectToURLSearchParams, urlSearchParamsToObject} from "@/Composables/helper.js";

// sorting handler section
export function useTableSort(props) {
  // example sort param ?sort=-created_at
  const defaultSorts = props.defaultSort.split("-");
  const sortKey = ref(defaultSorts[1] || defaultSorts[0]);
  const sortOrder = ref(defaultSorts.length > 1 ? "desc" : "asc");

  // function to sort column
  const sortColumn = (key) => {
    if (sortKey.value === key) {
      sortOrder.value = sortOrder.value === "asc" ? "desc" : "asc";
    } else {
      sortKey.value = key;
      sortOrder.value = "asc";
    }
    const sort =
      sortOrder.value === "desc" ? `-${sortKey.value}` : sortKey.value;

    const urlParams = new URLSearchParams(window.location.search);
    const mergedParams = mergeURLSearchParams(urlParams, objectToURLSearchParams({ sort: sort }));
    const paramsObject = urlSearchParamsToObject(mergedParams);

    router.get(
      `/auth/${props.module}`,
        paramsObject,
      { preserveState: true },
    );
  };

  return {
    sortKey,
    sortOrder,
    sortColumn,
  };
}
