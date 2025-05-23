import { FilterItem } from "../components/SearchAdvanceFilter";
import { FilterItem as FilterItemProto } from "../proto/genjs/utils/v1/utils_pb";

export const convertFilterToSearchAdvance = (filters: FilterItem[]) => {
    return filters.map(filter => {
        return {
            field: filter.filterCol,
            type: filter.filterType,
            filter: filter.value
        };
    });
}