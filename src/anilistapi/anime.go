package anilistapi

func SearchAnime(search string, page int, perPage int) (AniManga, error) {
	var anime AniManga

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
			media (id: $id, search: $search, type: ANIME, sort: POPULARITY_DESC) {
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

	err := post(query, variables, &anime)
	if err != nil {
		return anime, err
	}

	return anime, nil
}

func TopAnimeByScore(page int, perPage int) (AniManga, error) {
	var anime AniManga

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
			media (id: $id, type: ANIME, sort: SCORE_DESC) {
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

	err := post(query, variables, &anime)
	if err != nil {
		return anime, err
	}

	return anime, nil
}

func TopAnimeByPopularity(page int, perPage int) (AniManga, error) {
	var anime AniManga

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
			media (id: $id, type: ANIME, sort: POPULARITY_DESC) {
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

	err := post(query, variables, &anime)
	if err != nil {
		return anime, err
	}

	return anime, nil
}

func GetAnime(id int) (FullAnime, error) {
	var anime FullAnime

	query := `
	query ($id: Int) {
		Media (id: $id, type: ANIME) {
			id
			title {
				romaji
				english
				native
			}
			format
			episodes
			status
			season
			seasonYear
			averageScore
			popularity
			favourites
			source
			genres
			characters (role: MAIN) {
				edges {
					node {
						id
						name {
							full
						}
					}
					role
				}
			}
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

	err := post(query, variables, &anime)
	if err != nil {
		return anime, err
	}

	return anime, nil
}
