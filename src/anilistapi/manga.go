package anilistapi

func SearchManga(search string, page int, perPage int) (AniManga, error) {
	var manga AniManga

	checkPageValues(&page, &perPage)

	query := `
	query ($id: Int, $page: Int, $perPage: Int, $search: String) {
		Page (page: $page, perPage: $perPage) {
			pageInfo {
				total
				currentPage
				lastPage
				perPage
			}
			media (id: $id, search: $search, type: MANGA, sort: POPULARITY_DESC) {
				id
				title {
					romaji
					english
				}
				format
			}
		}
	}
	`

	variables := map[string]interface{}{
		"search":  search,
		"page":    page,
		"perPage": perPage,
	}

	err := post(query, variables, &manga)
	if err != nil {
		return manga, err
	}

	return manga, nil
}

func TopMangaByScore(page int, perPage int) (AniManga, error) {
	var manga AniManga

	checkPageValues(&page, &perPage)

	query := `
	query ($id: Int, $page: Int, $perPage: Int) {
		Page (page: $page, perPage: $perPage) {
			pageInfo {
				total
				currentPage
				lastPage
				perPage
			}
			media (id: $id, type: MANGA, sort: SCORE_DESC) {
				id
				title {
					romaji
					english
				}
				format
				averageScore
			}
		}
	}
	`

	variables := map[string]interface{}{
		"perPage": perPage,
		"page":    page,
	}

	err := post(query, variables, &manga)
	if err != nil {
		return manga, err
	}

	return manga, nil
}

func TopMangaByPopularity(page int, perPage int) (AniManga, error) {
	var manga AniManga

	checkPageValues(&page, &perPage)

	query := `
	query ($id: Int, $page: Int, $perPage: Int) {
		Page (page: $page, perPage: $perPage) {
			pageInfo {
				total
				currentPage
				lastPage
				perPage
			}
			media (id: $id, type: MANGA, sort: POPULARITY_DESC) {
				id
				title {
					romaji
				}
				format
				popularity
			}
		}
	}
	`

	variables := map[string]interface{}{
		"perPage": perPage,
		"page":    page,
	}

	err := post(query, variables, &manga)
	if err != nil {
		return manga, err
	}

	return manga, nil
}

func GetManga(id int) (FullManga, error) {
	var manga FullManga

	query := `
	query ($id: Int) {
		Media (id: $id, type: MANGA) {
			id
			title {
				romaji
				english
				native
			}
			format
			chapters
			volumes
			status
			startDate {
				day
				month
				year
			}
			averageScore
			popularity
			favourites
			source
			genres
			description (asHtml: false)
			coverImage {
				extraLarge
			}
		}
	}
	`

	variables := map[string]interface{}{
		"id": id,
	}

	err := post(query, variables, &manga)
	if err != nil {
		return manga, err
	}

	return manga, nil
}
